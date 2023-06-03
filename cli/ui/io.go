package ui

import (
	"fmt"
	"io"
)

type IO struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func (io *IO) Print(a ...any) (n int, err error) {
	return fmt.Fprint(io.Stdout, a...)
}

func (io *IO) Printf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(io.Stdout, format, a...)
}

func (io *IO) Println(a ...any) (n int, err error) {
	return fmt.Fprintln(io.Stdout, a...)
}

func (io *IO) Eprint(a ...any) (n int, err error) {
	return fmt.Fprint(io.Stderr, a...)
}

func (io *IO) Eprintf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(io.Stderr, format, a...)
}

func (io *IO) Eprintln(a ...any) (n int, err error) {
	return fmt.Fprintln(io.Stderr, a...)
}
