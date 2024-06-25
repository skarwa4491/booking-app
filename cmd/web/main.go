package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/skarwa4491/bookings/pkg/config"
	"github.com/skarwa4491/bookings/pkg/handlers"
	"github.com/skarwa4491/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is main application entry point
func main() {

	// change this to true when in production
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
	app.UseCache = false
	render.NewTemplate(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("starting application on port %s\n", portNumber)
	// starts web server on localhost 8080
	//_ = http.ListenAndServe(portNumber, nil)
	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}
	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal("unable to start server")
	}
}
