package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func runCmd(name string, args ...string) {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(name, args...)
	cmd.Dir = currentDir
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

func ensureDir(dir string) {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir = filepath.Join(currentDir, dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
