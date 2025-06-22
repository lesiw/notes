package main

import (
	"cmp"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"lesiw.io/flag"
)

var (
	flags = flag.NewSet(os.Stderr, "notes [-iw] [dir]")
	fInit = flags.Bool("i,init", "force creation of NOTES in dir")
	fVer  = flags.Bool("V,version", "print version")

	//go:embed version.txt
	versionfile string
	version     = strings.TrimRight(versionfile, "\n")

	cwd string
)

func main() {
	if err := run(); err != nil {
		if err.Error() != "" {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}

func run() (err error) {
	if err := flags.Parse(os.Args[1:]...); err != nil {
		return errors.New("")
	}
	if *fVer {
		fmt.Println(version)
		return nil
	}
	if cwd, err = os.Getwd(); err != nil {
		return fmt.Errorf("could not get current directory: %w", err)
	}
	dir, err := getOverlayDir()
	if err != nil {
		return fmt.Errorf("could not find dir from NOTESOVERLAY: %w", err)
	}
	if dir == "" {
		dir = cwd
	}
	if *fInit {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("could not create directory %s: %w", dir, err)
		}
		if err = os.WriteFile(notesFile(dir), []byte{}, 0644); err != nil {
			return fmt.Errorf("could not create NOTES file: %w", err)
		}
	}
	for {
		fi, err := os.Stat(notesFile(dir))
		if err == nil && !fi.IsDir() {
			break
		}
		if dir == "/" || dir == (filepath.VolumeName(dir)+"\\") {
			return errors.New("NOTES file not found")
		}
		dir = filepath.Join(dir, "..")
	}
	err = exec("/bin/sh", "-c", cmp.Or(
		os.Getenv("NOTESEDITOR"), os.Getenv("EDITOR"), "nano",
	)+" "+notesFile(dir))
	if err != nil {
		return fmt.Errorf("exec failed: %w", err)
	}
	return
}

func getOverlayDir() (dir string, err error) {
	for layer := range strings.SplitSeq(os.Getenv("NOTESOVERLAY"), "::") {
		top, btm, ok := strings.Cut(layer, ":")
		if !ok {
			continue
		}
		if !subDir(btm, cwd) {
			continue
		}
		rel, err := filepath.Rel(btm, cwd)
		if err != nil {
			continue
		}
		return filepath.Join(top, rel), nil
	}
	return
}

func subDir(parent, sub string) bool {
	parentpath, err := filepath.Abs(parent)
	if err != nil {
		return false
	}

	subpath, err := filepath.Abs(sub)
	if err != nil {
		return false
	}

	parentpath = filepath.Clean(parentpath) + string(filepath.Separator)
	subpath = filepath.Clean(subpath) + string(filepath.Separator)

	return strings.HasPrefix(subpath, parentpath)
}

func notesFile(dir string) string {
	return filepath.Join(dir, "NOTES")
}
