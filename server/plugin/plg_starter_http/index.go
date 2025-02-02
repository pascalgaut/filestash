package plg_starter_http

import (
	"fmt"
	"net/http"
	"time"

	. "github.com/pascalgaut/filestash/server/common"

	"github.com/gorilla/mux"
)

func init() {
	Hooks.Register.Starter(func(r *mux.Router) {
		Log.Info("[http] starting ...")
		port := Config.Get("general.port").Int()
		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: r,
		}
		go ensureAppHasBooted(
			fmt.Sprintf("http://127.0.0.1:%d%s", port, WithBase("/about")),
			fmt.Sprintf("[http] listening on :%d", port),
		)
		if err := srv.ListenAndServe(); err != nil {
			Log.Error("error: %v", err)
			return
		}
	})
}

func ensureAppHasBooted(address string, message string) {
	i := 0
	for {
		if i > 10 {
			Log.Warning("[http] didn't boot")
			break
		}
		time.Sleep(250 * time.Millisecond)
		res, err := http.Get(address)
		if err != nil {
			i += 1
			continue
		}
		res.Body.Close()
		if res.StatusCode != http.StatusOK {
			i += 1
			continue
		}
		Log.Info(message)
		break
	}
}
