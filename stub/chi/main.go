package main

import (
	"flag"
	"net/http"
	"time"

	// "PROJECT_NAME/html"

	// "github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	addr = flag.String("addr", ":4000", "HTTP server address")
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(time.Second * 30))
	r.Use(middleware.NoCache)

	// r.Get("/", templ.Handler(html.HomePage()).ServeHTTP)

	s := http.Server{
		Addr:    *addr,
		Handler: r,
	}

	s.ListenAndServe()
}
