//go:build windows

package main

import (
	"os"
	osexec "os/exec"
)

func exec(args ...string) error {
	cmd := osexec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
