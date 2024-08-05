package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/MeherKandukuri/reservationSystem_Go/internal/config"
	"github.com/MeherKandukuri/reservationSystem_Go/internal/models"

	"github.com/justinas/nosurf"
)

var app *config.AppConfig
var pathToTemplates string = "./templates"

// AddDefaultData contains Data which will be added to data sent to templates
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func NewTemplates(a *config.AppConfig) {
	app = a
}

// renderTemplate serves as a wrapper and renders a layout and a template from folder /templates to desired writer
func RenderTemplate(w http.ResponseWriter, r *http.Request, tpml string, td *models.TemplateData) error {
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get the right template from cache
	t, ok := tc[tpml]
	if !ok {
		return errors.New("template not in cache for some reason")
	}

	// render that template
	buf := new(bytes.Buffer)

	// we can add Default Data here:
	td = AddDefaultData(td, r)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}
	// render that template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	return nil
}

// CreateTemplateCache creates a map and stores the templates in for cachine
func CreateTemplateCache() (map[string]*template.Template, error) {

	theCache := map[string]*template.Template{}

	// get all the available files *-page.tpml from folder ./templates
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

/*
-- Templating with Dynamic Caching
// template cache (tc) for holding templates
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc[t]
	var count int
	if !inMap {
		// need to create the template
		count++
		log.Println("Creating template and adding to cache")
		log.Println("count:",count)
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have template in the cache
		log.Println("using cached template")
	}
	// assigning the template to template variable
	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base-layout.html",
	}
	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to cache (map)
	tc[t] = tmpl
	return nil

}
*/

/*
Templating without Caching
func RenderTemplateTemp(w http.ResponseWriter, html string) {

	parsedTemplate, err := template.ParseFiles("./templates/"+html, "./templates/base-layout.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Unable to parse template", http.StatusInternalServerError)
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}

}
*/
