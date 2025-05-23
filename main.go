package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	switch os.Args[1] {
	case "chi":
		chi()
		return

	case "help", "-help", "--help", "-h":
		usage()
		return
	}
}
