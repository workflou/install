package main

import (
	"flag"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var (
	addr = flag.String("addr", ":4000", "HTTP server address")
)

func main() {
	r := chi.NewRouter()

	s := http.Server{
		Addr:    *addr,
		Handler: r,
	}

	s.ListenAndServe()
}
