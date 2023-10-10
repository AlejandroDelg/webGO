package main

import (
	"net/http"

	"github.com/AlejandroDelg/webgo/pkg/config"
	"github.com/AlejandroDelg/webgo/pkg/handlers"

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
	mux.Get("/weapons", handlers.Repo.Weapons)
	mux.Get("/quests", handlers.Repo.Quests)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation-quest", handlers.Repo.MakeReservationQuest)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}
