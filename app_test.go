package gotmpl2html_test

import (
	"bytes"
	"testing"

	. "github.com/grape80/gotmpl2html"
)

func TestApp_Run(t *testing.T) {
	tests := map[string]struct {
		exitCodeExpected int
		args             []string
	}{
		"no_args": {EXIT_FAILURE, []string{}},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(tt *testing.T) {
			app := New()
			app.Stdin = new(bytes.Buffer)
			app.Stdout = new(bytes.Buffer)
			app.Stderr = new(bytes.Buffer)
			app.Name = "gotmpl2html"
			app.Args = test.args

			exitCode := app.Run()
			if exitCode != test.exitCodeExpected {
				t.Errorf("Run() = %v; want %v", exitCode, test.exitCodeExpected)
			}
		})
	}
}
