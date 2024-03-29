package main

import (
	"encoding/gob"
	"fmt"
	"github.com/AlejandroDelg/webgo/helpers"
	"github.com/AlejandroDelg/webgo/internal/config"
	"github.com/AlejandroDelg/webgo/internal/driver"
	"github.com/AlejandroDelg/webgo/internal/handlers"
	"github.com/AlejandroDelg/webgo/internal/models"
	"github.com/AlejandroDelg/webgo/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main function
func main() {
	err, db := run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Staring application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err2 := srv.ListenAndServe()

	if err2 != nil {
		fmt.Println("Error in Server: ", err2)
	}

	defer db.SQL.Close()
}

func run() (error, *driver.DB) {

	gob.Register(models.Reservation{})
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO   ", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR    ", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	log.Println("Connecting to database ...")
	// create database
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=webGo user=postgres password=root")

	if err != nil{
		log.Fatal("ERROR: ", err)
	}

	log.Println("Connected to database !!!")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err, nil
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

	render.NewTemplates(&app)

	helpers.NewHelpers(&app)

	repo := handlers.NewRepo(&app, db)

	handlers.NewHandlers(repo)

	handlers.GetMonsters(allMonsters)
	return nil, db
}
