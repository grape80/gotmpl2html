package conv

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type Converter struct {
	Basedir  string
	Gotmpl   string
	Partials []string
}

func (c *Converter) Run() (html string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s %v\n", ERR_CONV_RUNTIME, r)
		}
	}()

	if err = c.parseTemplate(); err != nil {
		return "", err
	}

	return c.execute()
}

const (
	partialPrefix = `{{template "`
	partialSuffix = `}}`
)

func (c *Converter) parseTemplate() error {
	f, err := os.Open(c.Gotmpl)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := strings.TrimSpace(s.Text()) // to clean
		if strings.HasPrefix(l, partialPrefix) && strings.HasSuffix(l, partialSuffix) {
			p := filepath.Join(c.Basedir, strings.Split(l, `"`)[1]) // prefix”path"suffix -> [prefix path suffix]
			c.Partials = append(c.Partials, p)
		}
	}

	return nil
}

func (c *Converter) execute() (string, error) {
	files := []string{c.Gotmpl}
	files = append(files, c.Partials...)

	t := template.Must(template.ParseFiles(files...))

	var buf bytes.Buffer
	if err := t.Execute(&buf, nil); err != nil {
		return "", err
	}

	return buf.String(), nil
}
