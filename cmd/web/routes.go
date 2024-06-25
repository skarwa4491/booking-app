package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/skarwa4491/bookings/pkg/handlers"
)

// routes using pat router
//func routes(app *config.AppConfig) http.Handler {
//
//	mux := pat.New()
//
//	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
//	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
//	return mux
//}

// using chi router

func routes() http.Handler {

	mux := chi.NewRouter()

	// middleware here: middle ware use to process request and perform some operation on request before executing
	mux.Use(WriteToConsole)
	mux.Use(SessionLoad)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// creating a file server
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
