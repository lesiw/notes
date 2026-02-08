//go:build !windows

package main

import (
	"fmt"
	"os"
	osexec "os/exec"
	"syscall"
)

func exec(args ...string) error {
	path, err := osexec.LookPath(args[0])
	if err != nil {
		return fmt.Errorf("lookpath: %w", err)
	}
	return syscall.Exec(path, args, os.Environ())
}
