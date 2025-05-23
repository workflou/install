package main

import (
	"os"
	"os/exec"
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
