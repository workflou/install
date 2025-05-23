package main

func chi() {
	copyStub("chi/main.go", "main.go")
	copyStub("chi/gitignore", ".gitignore")

	runCmd("go", "get", "github.com/go-chi/chi/v5")
	runCmd("go", "mod", "tidy")
}
