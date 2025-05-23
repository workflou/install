package main

func templ() {
	ensureDir("html")

	copyStub("templ/html/home_page.templ", "html/home_page.templ")

	runCmd("go", "get", "-tool", "github.com/a-h/templ/cmd/templ@latest")
	runCmd("go", "get", "github.com/a-h/templ/cmd/templ@latest")
	runCmd("go", "mod", "tidy")
}
