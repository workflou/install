package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/workflou/install/stub"
)

func copyStub(src string, dst string) {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(filepath.Join(currentDir, dst)); err == nil {
		fmt.Printf("%s already exists in the current directory.\n", dst)
		return
	}

	file, err := fs.ReadFile(stub.FS, src)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(currentDir, dst), file, 0644)
	if err != nil {
		panic(err)
	}
}
