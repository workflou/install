package main

func tailwind() {
	ensureDir("static")
	ensureDir("static/css")

	copyStub("static/css/main.css", "static/css/main.css")
}
