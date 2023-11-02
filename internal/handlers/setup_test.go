package handlers

import (
	"encoding/gob"
	"fmt"
	"github.com/AlejandroDelg/webgo/internal/config"
	"github.com/AlejandroDelg/webgo/internal/driver"
	"github.com/AlejandroDelg/webgo/internal/models"
	"github.com/AlejandroDelg/webgo/internal/render"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var pathToTemplates = "./../../templates"

var functions = template.FuncMap{}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true, Path: "/", Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)

}

var app config.AppConfig
var session *scs.SessionManager

func getRoutes() http.Handler {

	gob.Register(models.Reservation{})
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	allMonsters := []*models.Monster{}
	allQuests := []*models.Quest{}

	fatalis := models.Fatalis()
	lagiacrus := models.Lagiacrus()
	questFatalis := models.QuestFatalis()
	fatalis.Quests = append(fatalis.Quests, questFatalis)

	allMonsters = append(allMonsters, lagiacrus)
	allMonsters = append(allMonsters, fatalis)
	allQuests = append(allQuests, &questFatalis)

	app.TemplateCache = tc
	app.UseCache = true

	infoLog := log.New(os.Stdout, "INFO   ", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR    ", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	render.NewTemplates(&app)

	log.Println("Connecting to database ...")

	// create database
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=webGo user=postgres password=root")

	if err != nil{
		log.Fatal("ERROR: ", err)
	}

	log.Println("Connected to database !!!")

	
	repo := NewRepo(&app, db)

	NewHandlers(repo)

	GetMonsters(allMonsters)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/monsters", Repo.Monsters)
	mux.Get("/monsters/{name}", Repo.Monster)
	mux.Get("/weapons", Repo.Weapons)
	mux.Get("/quests", Repo.Quests)

	mux.Get("/contact", Repo.Contact)

	mux.Get("/make-reservation-quest", Repo.MakeReservationQuest)
	mux.Post("/make-reservation-quest", Repo.PostMakeReservationQuest)

	mux.Get("/make-reservation", Repo.MakeReservation)
	mux.Post("/make-reservation", Repo.PostMakeReservation)

	mux.Get("/reservation-summary", Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}

// CreateTemplateCache creates a template cache as a map
func CreateTestTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}
	println(fmt.Sprintf("%s/*.html", pathToTemplates))
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.html", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {

		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
