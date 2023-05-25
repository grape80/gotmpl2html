package conv

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/exp/maps"
)

type Converter struct {
	Basedir    string
	Gotmpl     string
	partials   []string
	data       map[int][]string
	dataObject map[string]any
}

func NewConverter() *Converter {
	return &Converter{
		data:       make(map[int][]string),
		dataObject: make(map[string]any),
	}
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

	if err = c.createDataObject(); err != nil {
		return "", err
	}

	return c.execute()
}

const (
	partialPrefix = `{{template "`
	partialSuffix = `}}`

	dataPrefix = `{{data "`
	dataSuffix = `}}`
)

func (c *Converter) parseTemplate() error {
	depth := 0

	var fn func(string) error
	fn = func(path string) error {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		s := bufio.NewScanner(f)
		for s.Scan() {
			l := strings.TrimSpace(s.Text()) // to clean
			switch {
			case strings.HasPrefix(l, partialPrefix) && strings.HasSuffix(l, partialSuffix):
				p := filepath.Join(c.Basedir, strings.Split(l, `"`)[1]) // prefix”path"suffix -> [prefix path suffix]
				c.partials = append(c.partials, p)

				depth++
				fn(p)
			case strings.HasPrefix(l, dataPrefix) && strings.HasSuffix(l, dataSuffix):
				d := filepath.Join(c.Basedir, strings.Split(l, `"`)[1]) // prefix”path"suffix -> [prefix path suffix]
				c.data[depth] = append(c.data[depth], d)
			}
		}

		depth--

		return nil
	}

	return fn(c.Gotmpl)
}

func (c *Converter) createDataObject() (err error) {
	for l := len(c.data); l > 0; l-- {
		for _, d := range c.data[l-1] {
			var b []byte
			if b, err = os.ReadFile(d); err != nil {
				return err
			}

			var m map[string]any
			if err = json.Unmarshal(b, &m); err != nil {
				return err
			}
			maps.Copy(c.dataObject, m)
		}
	}

	return nil
}

func (c *Converter) execute() (string, error) {
	funcs := map[string]any{
		"data": func(s string) string {
			// no operation
			return ""
		},
	}

	files := []string{c.Gotmpl}
	files = append(files, c.partials...)

	t := template.Must(template.New(filepath.Base(c.Gotmpl)).Funcs(funcs).ParseFiles(files...))

	var buf bytes.Buffer
	if err := t.Execute(&buf, c.dataObject); err != nil {
		return "", err
	}

	return buf.String(), nil
}
