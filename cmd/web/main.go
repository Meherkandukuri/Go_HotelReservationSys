// package main

// import (
// 	"fmt"
// 	"log"
// 	"github.com/MeherKandukuri/reservationSystem_Go/pkg/config"
// 	"github.com/MeherKandukuri/reservationSystem_Go/pkg/handlers"
// 	"github.com/MeherKandukuri/reservationSystem_Go/pkg/render"
// 	"net/http"
// 	"time"

// 	"github.com/alexedwards/scs/v2"
// )

// const portNumber = ":8080"

// var app config.AppConfig
// var session *scs.SessionManager

// func main() {

// 	// dont forget to change to true in Production!
// 	app.InProduction = false

// 	session = scs.New()
// 	session.Lifetime = 24 * time.Hour
// 	session.Cookie.Persist = true
// 	session.Cookie.SameSite = http.SameSiteLaxMode
// 	session.Cookie.Secure = app.InProduction
// 	app.Session = session

// 	tc, err := render.CreateTemplateCache()
// 	if err != nil {
// 		log.Fatalln("error creating template cache", err)
// 	}
// 	app.TemplateCache = tc
// 	app.UseCache = false

// 	repo := handlers.NewRepo(&app)
// 	handlers.NewHandlers(repo)

// 	render.NewTemplates(&app)

// 	fmt.Println("We are starting in port number:", portNumber)

// 	srv := &http.Server{
// 		Addr:    portNumber,
// 		Handler: routes(&app),
// 	}

//		err = srv.ListenAndServe()
//		if err != nil {
//			log.Fatalln(err)
//		}
//	}
//
// main.go
package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"github.com/MeherKandukuri/reservationSystem_Go/internal/config"
	"github.com/MeherKandukuri/reservationSystem_Go/internal/handlers"
	"github.com/MeherKandukuri/reservationSystem_Go/internal/models"
	"github.com/MeherKandukuri/reservationSystem_Go/internal/render"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("We are starting on port number:", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	// Data to be available in the session
	gob.Register(models.Reservation{})

	// don't forget to change to true in Production!
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("error creating template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	return nil
}
