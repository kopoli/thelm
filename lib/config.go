package thelm

import (
	"github.com/casimir/xdg-go"
	util "github.com/kopoli/go-util"
)

// DefaultConfigFile gets the default configuration file name based on given
// Options
func DefaultConfigFile(opts util.Options) string {
	app := xdg.App{Name: opts.Get("application-name", "gogr")}
	return app.ConfigPath(opts.Get("configuration-file", "config.json"))
}
