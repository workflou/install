package main

func chi() {
	copyStub("chi/main.go", "main.go")

	runCmd("go", "get", "github.com/go-chi/chi/v5")
	runCmd("go", "mod", "tidy")
}
