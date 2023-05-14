package commands

import "github.com/grape80/gotmpl2html/cli/ui"

type Help struct {
	*ui.IO
	Usage string
}

func (c *Help) Execute() error {
	_, err := c.Print(c.Usage)
	return err
}
