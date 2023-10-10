package main

import (
	"fmt"
	"github.com/AlejandroDelg/webgo/internal/config"
	"github.com/AlejandroDelg/webgo/internal/handlers"
	"github.com/AlejandroDelg/webgo/internal/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	fmt.Printf("Staring application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err2 := srv.ListenAndServe()

	if err2 != nil {
		fmt.Println("Error in Server: ", err2)
	}
}
