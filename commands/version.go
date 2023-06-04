package commands

import "github.com/grape80/cli/ui"

type Version struct {
	*ui.IO
	Version string
}

func (c *Version) Execute() error {
	_, err := c.Println(c.Version)
	return err
}
