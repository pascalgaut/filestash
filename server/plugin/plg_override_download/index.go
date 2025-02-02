package plg_override_download

import (
	"embed"

	. "github.com/pascalgaut/filestash/server/common"
)

//go:embed assets/*
var STATIC embed.FS

func init() {
	Hooks.Register.StaticPatch(STATIC)
}
