package plg_starter_tor

import (
	"context"
	"github.com/cretz/bine/tor"
	"github.com/gorilla/mux"
	. "github.com/pascalgaut/filestash/server/common"
	"net/http"
	"os"
	"time"
)

var (
	enable_plugin func() bool
	tor_url       func() string
)

func init() {
	enable_plugin = func() bool {
		return Config.Get("features.server.tor_enable").Schema(func(f *FormElement) *FormElement {
			if f == nil {
				f = &FormElement{}
			}
			f.Default = false
			f.Name = "tor_enable"
			f.Type = "enable"
			f.Target = []string{"tor_url"}
			f.Description = "Enable/Disable tor server"
			f.Placeholder = "Default: false"
			return f
		}).Bool()
	}
	tor_url = func() string {
		return Config.Get("features.server.tor_url").Schema(func(f *FormElement) *FormElement {
			if f == nil {
				f = &FormElement{}
			}
			f.Id = "tor_url"
			f.Name = "tor_url"
			f.Type = "text"
			f.Target = []string{}
			f.Description = "Your onion site"
			f.ReadOnly = true
			f.Placeholder = "LOADING... Refresh the page in a few seconds"
			return f
		}).String()
	}

	Hooks.Register.Onload(func() {
		tor_url()
		enable_plugin()
	})
	Hooks.Register.Starter(func(r *mux.Router) {
		torPath := GetAbsolutePath(CERT_PATH, "tor")
		os.MkdirAll(torPath, os.ModePerm)

		if enable_plugin() == false {
			startTor := false
			onChange := Config.ListenForChange()
			for {
				select {
				case <-onChange.Listener:
					startTor = enable_plugin()
				}
				if startTor == true {
					break
				}
			}
			Config.UnlistenForChange(onChange)
		}

		Log.Info("[tor] starting ...")
		t, err := tor.Start(nil, &tor.StartConf{
			DataDir: torPath,
		})
		if err != nil {
			Log.Error("[tor] Unable to start Tor: %v", err)
			return
		}
		defer t.Close()
		listenCtx, listenCancel := context.WithTimeout(context.Background(), 3*time.Minute)
		defer listenCancel()
		onion, err := t.Listen(listenCtx, &tor.ListenConf{Version3: true, RemotePorts: []int{80}})
		if err != nil {
			Log.Error("[tor] Unable to create onion service: %v", err)
			return
		}
		defer onion.Close()

		srv := &http.Server{
			Handler: r,
		}
		Log.Info("[tor] started http://%s.onion\n", onion.ID)
		Config.Get("features.server.tor_url").Set("http://" + onion.ID + ".onion")
		srv.Serve(onion)
	})
}
