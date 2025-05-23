package main

func air() {
	copyStub("air/air.toml", ".air.toml")

	runCmd("go", "get", "-tool", "github.com/air-verse/air@latest")
	runCmd("go", "mod", "tidy")
}
