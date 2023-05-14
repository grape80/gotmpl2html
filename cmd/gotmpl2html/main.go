package main

import (
	"os"

	"github.com/grape80/gotmpl2html"
)

var _version string

func main() {
	app := gotmpl2html.New()
	app.Version = _version

	os.Exit(app.Run())
}
