package main

func htmx() {
	ensureDir("static")
	ensureDir("static/js")

	copyStub("static/embed.go", "static/embed.go")
	copyStub("static/js/htmx@2.0.4.min.js", "static/js/htmx@2.0.4.min.js")
}
