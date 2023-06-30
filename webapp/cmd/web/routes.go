package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	// Register middleware

	mux.Use(middleware.Recoverer)
	mux.Use(app.addIpToCtx)

	// Register routes
	mux.Get("/", app.Home)
	mux.Post("/login", app.Login)
	// Static assets
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
