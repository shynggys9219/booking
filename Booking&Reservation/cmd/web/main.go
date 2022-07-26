package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/shynggys9219/goBookingProject/config"
	"github.com/shynggys9219/goBookingProject/pkg/handlers"
	"github.com/shynggys9219/goBookingProject/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is main application function
func main() {

	// change this to true when in production
	app.InProduction = false

	// create a new session
	session = scs.New()

	session.Lifetime = 24 * time.Hour
	// cookies won't disappear after closing browser if Persist is set true
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	// for development purposes set false only
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	// template cache stored in struct
	app.TemplateCach = tc

	app.UseCache = true // usually false in development mode, in production is true
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//// handlers.Repo.* is saying that these handlers have receivers
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting app on port %s", portNumber)
	//_ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
