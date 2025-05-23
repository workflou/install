package main

import "fmt"

func usage() {
	fmt.Println("Usage: go tool install <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  chi    Add chi router")
	fmt.Println("Options:")
	fmt.Println("  -h, --help    Show this help message")
}
