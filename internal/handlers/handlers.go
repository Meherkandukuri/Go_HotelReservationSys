package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/config"
	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/driver"
	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/forms"
	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/helpers"
	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/models"
	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/render"
	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/repository"
	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/repository/dbrepo"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// Repo the repository used by the handlers
var Repo *Repository

// New Repo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostGresRepo(db.SQL, a),
	}
}

// NeweHandlers sets the repositry for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home-page.html", &models.TemplateData{})
}

// About is the handler for about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// send the result or any prepared data to the template
	render.Template(w, r, "about-page.html", &models.TemplateData{})
}

// Contact is the handler for contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	// send the result or any prepared data to the template
	render.Template(w, r, "contact-page.html", &models.TemplateData{})
}

// Eremite is the handler for Eremite page
func (m *Repository) Eremite(w http.ResponseWriter, r *http.Request) {
	// send the result or any prepared data to the template
	render.Template(w, r, "eremite-page.html", &models.TemplateData{})
}

// Couple is the handler for the couple page
func (m *Repository) Couple(w http.ResponseWriter, r *http.Request) {
	// send the result or any prepared data to the template
	render.Template(w, r, "couple-page.html", &models.TemplateData{})
}

// Family is the handler for the family page
func (m *Repository) Family(w http.ResponseWriter, r *http.Request) {
	// send the result or any prepared data to the template
	render.Template(w, r, "family-page.html", &models.TemplateData{})
}

// Reservation is the handler for reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	// send the result or any prepared data to the template
	render.Template(w, r, "check-availability-page.html", &models.TemplateData{})
}

// PostReservation is the handler for the reservation page and POST requests
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("startingDate")
	end := r.Form.Get("endingDate")
	layout := "2006-01-02"

	startDate, err := time.Parse(layout, start)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	bungalows, err := m.DB.SearchAvailabilityByDatesForAllBungalows(startDate, endDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	if len(bungalows) == 0 {
		m.App.Session.Put(r.Context(), "error", ":( No holiday home is available at that time.")
		http.Redirect(w, r, "/reservation", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["bungalows"] = bungalows

	res := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}

	m.App.Session.Put(r.Context(), "reservation", res)

	render.Template(w, r, "choose-bungalow-page.html", &models.TemplateData{
		Data: data,
	})

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
		helpers.ServerError(w, err)
		return
	}

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
	render.Template(w, r, "make-reservation-page.html",
		&models.TemplateData{
			Form: forms.New(nil),
			Data: data,
		})
}

// PostMakeReservation is the Post request handler for the make-reservation page
func (m *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")

	layout := "2006-01-02"

	startDate, err := time.Parse(layout, sd)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	endDate, err := time.Parse(layout, ed)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	bungalowID, err := strconv.Atoi(r.Form.Get("bungalow_id"))
	reservation := models.Reservation{
		FullName:   r.Form.Get("full_name"),
		Email:      r.Form.Get("email"),
		Phone:      r.Form.Get("phone"),
		StartDate:  startDate,
		EndDate:    endDate,
		BungalowID: bungalowID,
	}

	form := forms.New(r.PostForm)

	// form.Has("full_name", r)
	form.Required("full_name", "email")
	form.MinLength("full_name", 2)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.Template(w, r, "make-reservation-page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	newReservationID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	restriction := models.BungalowRestriction{
		StartDate:     startDate,
		EndDate:       endDate,
		BungalowID:    bungalowID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = m.DB.InsertBungalowRestriction(restriction)
	if err != nil {
		helpers.ServerError(w, err)
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
		m.App.ErrorLog.Println("Could not get item from session")
		m.App.Session.Put(r.Context(), "error", "No reservation data in this session is available.")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.Template(w, r, "reservation-overview-page.html", &models.TemplateData{
		Data: data,
	})
}
