package commands

import (
	"errors"
	"flag"
	"io"
	"strings"

	"github.com/grape80/gotmpl2html/cli/ui"
	"github.com/grape80/gotmpl2html/conv"
	"github.com/grape80/gotmpl2html/msgs"
)

type Run struct {
	*ui.IO
	basedir string
	gotmpl  string
}

func (c *Run) Execute() error {
	conv := conv.NewConverter()
	conv.Basedir = c.basedir
	conv.Gotmpl = c.gotmpl

	html, err := conv.Run()
	if err != nil {
		return err
	}

	_, err = c.Print(html)

	return err
}

func (c *Run) ParseArgs(args []string) error {
	f := flag.NewFlagSet("run", flag.ContinueOnError)
	f.StringVar(&c.basedir, "basedir", "./", "")
	f.SetOutput(io.Discard)

	if err := f.Parse(args); err != nil {
		m := strings.Replace(err.Error(), " -", " ", 1) // -option -> option
		return errors.New(m)
	}

	gotmpls := f.Args()
	switch len(gotmpls) {
	case 1:
		c.gotmpl = gotmpls[0]
		return nil
	case 0:
		return errors.New(msgs.ERR_NO_GOTMPL)
	default:
		return errors.New(msgs.ERR_MULTI_GOTMPLS)
	}
}
