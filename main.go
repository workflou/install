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

	case "air":
		air()

	case "makefile":
		makefile()

	case "sqlc":
		sqlc()

	case "templ":
		templ()

	case "htmx":
		htmx()

	case "help", "-help", "--help", "-h":
		usage()

	case "all":
		chi()
		air()
		sqlc()
		templ()
		htmx()
		makefile()

	default:
		usage()
	}
}
