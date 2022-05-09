package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jerevick83/hotel_bookings/pkg/config"
	"github.com/jerevick83/hotel_bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(NoSurf)
	r.Use(SessionLoad)
	r.HandleFunc("/", handlers.Repo.Home)
	r.HandleFunc("/about", handlers.Repo.About)
	return r
}
