package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/config"
	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/driver"
	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/handlers"
	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/helpers"
	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/models"
	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer close(app.MailChan)

	fmt.Println("Starting E-mail listener")
	listenForMail()


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

func run() (*driver.DB, error) {
	// Data to be available in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Bungalow{})
	gob.Register(models.BungalowRestriction{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	// don't forget to change to true in Production!
	app.InProduction = false

	infoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "[Error]\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=GoHotelReservationSystem user=postgres password=")
	if err != nil {
		log.Fatal("No connection to database! Terminating ...")
	}
	// defer db.SQL.Close()
	log.Println("Successfully connected to database.")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("error creating template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)

	helpers.NewHelpers(&app)

	return db, nil
}
