package ui_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/grape80/gotmpl2html/cli/ui"
)

func TestIO_Print(t *testing.T) {
	stdout := new(bytes.Buffer)
	io := &ui.IO{
		os.Stdin,
		stdout,
		os.Stderr,
	}

	input := "Output to stdout"
	io.Print(input)

	got := stdout.String()
	want := input
	if got != want {
		t.Errorf("Print() = %v; want %v", got, want)
	}
}

func TestIO_Printf(t *testing.T) {
	stdout := new(bytes.Buffer)
	io := &ui.IO{
		os.Stdin,
		stdout,
		os.Stderr,
	}

	input := "Output to stdout"
	io.Printf("%s", input)

	got := stdout.String()
	want := input
	if got != want {
		t.Errorf("Printf() = %v; want %v", got, want)
	}
}

func TestIO_Println(t *testing.T) {
	stdout := new(bytes.Buffer)
	io := &ui.IO{
		os.Stdin,
		stdout,
		os.Stderr,
	}

	input := "Output to stdout"
	io.Println(input)

	got := stdout.String()
	want := input + "\n"
	if got != want {
		t.Errorf("Println() = %v; want %v", got, want)
	}
}

func TestIO_Eprint(t *testing.T) {
	stderr := new(bytes.Buffer)
	io := &ui.IO{
		os.Stdin,
		os.Stdout,
		stderr,
	}

	input := "Output to stderr"
	io.Eprint(input)

	got := stderr.String()
	want := input
	if got != want {
		t.Errorf("Eprint() = %v; want %v", got, want)
	}
}

func TestIO_Eprintf(t *testing.T) {
	stderr := new(bytes.Buffer)
	io := &ui.IO{
		os.Stdin,
		os.Stdout,
		stderr,
	}

	input := "Output to stderr"
	io.Eprintf("%s", input)

	got := stderr.String()
	want := input
	if got != want {
		t.Errorf("Eprintf() = %v; want %v", got, want)
	}
}

func TestIO_Eprintln(t *testing.T) {
	stderr := new(bytes.Buffer)
	io := &ui.IO{
		os.Stdin,
		os.Stdout,
		stderr,
	}

	input := "Output to stderr"
	io.Eprintln(input)

	got := stderr.String()
	want := input + "\n"
	if got != want {
		t.Errorf("Eprintln() = %v; want %v", got, want)
	}
}
