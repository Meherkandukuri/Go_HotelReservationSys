package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/MeherKandukuri/reservationSystem_Go/internal/config"
	"github.com/MeherKandukuri/reservationSystem_Go/internal/models"
	"github.com/MeherKandukuri/reservationSystem_Go/internal/render"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"

func getRoutes() http.Handler {
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

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("error creating template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := NewRepo(&app)
	NewHandlers(repo)

	render.NewTemplates(&app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/contact", Repo.Contact)
	mux.Get("/eremite", Repo.Eremite)
	mux.Get("/couple", Repo.Couple)
	mux.Get("/family", Repo.Family)
	mux.Get("/reservation", Repo.Reservation)
	mux.Post("/reservation", Repo.PostReservation)
	mux.Post("/reservation-json", Repo.ReservationJSON)
	mux.Get("/make-reservation", Repo.MakeReservation)
	mux.Post("/make-reservation", Repo.PostMakeReservation)
	mux.Get("/reservation-overview", Repo.ReservationOverview)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

// NoSurf serves as a CSRF protection middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads saves session data for each request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// CreateTemplateCache creates a map and stores the templates in for cachine
func CreateTestTemplateCache() (map[string]*template.Template, error) {

	theCache := map[string]*template.Template{}

	// get all the available files *-page.tpml from folder templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*-page.html", pathToTemplates))

	if err != nil {
		return theCache, err
	}

	// range through the slice of *-page.tpml
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return theCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*-layout.html", pathToTemplates))
		if err != nil {
			return theCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*-layout.html", pathToTemplates))
			if err != nil {
				return theCache, err
			}
		}

		theCache[name] = ts
	}

	return theCache, nil

}
