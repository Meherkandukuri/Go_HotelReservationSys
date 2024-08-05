package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MeherKandukuri/reservationSystem_Go/internal/config"
	"github.com/MeherKandukuri/reservationSystem_Go/internal/forms"
	"github.com/MeherKandukuri/reservationSystem_Go/internal/models"
	"github.com/MeherKandukuri/reservationSystem_Go/internal/render"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// New Repo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NeweHandlers sets the repositry for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, r, "home-page.html", &models.TemplateData{})
}

// About is the handler for about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// send the result or any prepared data to the template
	render.RenderTemplate(w, r, "about-page.html", &models.TemplateData{})
}

// Contact is the handler for contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	// send the result or any prepared data to the template
	render.RenderTemplate(w, r, "contact-page.html", &models.TemplateData{})
}

// Eremite is the handler for Eremite page
func (m *Repository) Eremite(w http.ResponseWriter, r *http.Request) {
	// send the result or any prepared data to the template
	render.RenderTemplate(w, r, "eremite-page.html", &models.TemplateData{})
}

// Couple is the handler for the couple page
func (m *Repository) Couple(w http.ResponseWriter, r *http.Request) {
	// send the result or any prepared data to the template
	render.RenderTemplate(w, r, "couple-page.html", &models.TemplateData{})
}

// Family is the handler for the family page
func (m *Repository) Family(w http.ResponseWriter, r *http.Request) {
	// send the result or any prepared data to the template
	render.RenderTemplate(w, r, "family-page.html", &models.TemplateData{})
}

// Reservation is the handler for reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	// send the result or any prepared data to the template
	render.RenderTemplate(w, r, "check-availability-page.html", &models.TemplateData{})
}

// PostReservation is the handler for the reservation page and POST requests
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("startingDate")
	end := r.Form.Get("endingDate")

	w.Write([]byte(fmt.Sprintf("Arrival date value is set to %s, departure date is set to %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// ReservationJson is the handler for reservation-json route and return JSON
func (m *Repository) ReservationJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "It's available!",
	}

	output, err := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(output))
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// MakeReservation is the handler for the make-reservation page
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {

	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	/* we can write above three lines of code into inline statement one as follows:
	data := map[string]{}interface{
		"reservation":models.Reservation{},
	}
	*/
	// send the result or any prepared data to the template
	render.RenderTemplate(w, r, "make-reservation-page.html",
		&models.TemplateData{
			Form: forms.New(nil),
			Data: data,
		})
}

// PostMakeReservation is the Post request handler for the make-reservation page
func (m *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		Name:  r.Form.Get("full_name"),
		Email: r.Form.Get("email"),
		Phone: r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	// form.Has("full_name", r)
	form.Required("full_name", "email")
	form.MinLength("full_name", 2)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation-page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-overview", http.StatusSeeOther)
}

// ReservationOverview displays the reservation summary page
func (m *Repository) ReservationOverview(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("Could  not get item from session.")
		m.App.Session.Put(r.Context(), "error", "No reservation data in this session is available.")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.RenderTemplate(w, r, "reservation-overview-page.html", &models.TemplateData{
		Data: data,
	})
}
