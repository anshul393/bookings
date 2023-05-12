package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	config "github.com/anshul393/pkg/config"
	handler "github.com/anshul393/pkg/handlers"
	"github.com/anshul393/pkg/render"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // Browser will keep the session information even if we close the browser
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create the template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false // In Development Environment

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)

	mux := routes(&app)

	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	fmt.Println("Server is ready and start listening at", server.Addr)
	server.ListenAndServe()
}
