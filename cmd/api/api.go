package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wafiqpuyol/GO-Social/internal/store"
)

// NOTE => This app struct is our dependency
type application struct {
	config config
	store  store.Storage
}
type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr               string
	maxOpenConnections int
	maxIdleConnections int
	maxIdleTime        string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", app.health)
		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPost)
		})
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	server := http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}
	return server.ListenAndServe()
}
