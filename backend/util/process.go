package util

import (
	"io"
	"os/exec"
)

type ProcessAttr struct {
	Args   []string
	Env    []string
	Dir    string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func StartProcess(name string, attr *ProcessAttr) *exec.Cmd {
	p := exec.Command(name, attr.Args...)
	p.Env = attr.Env
	p.Dir = attr.Dir
	p.Stdin = attr.Stdin
	p.Stdout = attr.Stdout
	p.Stderr = attr.Stderr
	if err := p.Start(); err != nil {
		panic(err)
	}
	return p
}
