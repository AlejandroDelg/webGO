package main

import (
	"github.com/AlejandroDelg/webgo/internal/config"
	"github.com/AlejandroDelg/webgo/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/monsters", handlers.Repo.Monsters)
	mux.Get("/monsters/{name}", handlers.Repo.Monster)
	mux.Get("/weapons", handlers.Repo.Weapons)
	mux.Get("/quests", handlers.Repo.Quests)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation-quest", handlers.Repo.MakeReservationQuest)
	mux.Post("/make-reservation-quest", handlers.Repo.PostMakeReservationQuest)

	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Post("/make-reservation", handlers.Repo.PostMakeReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}
