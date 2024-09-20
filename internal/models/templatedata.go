package models

import "github.com/MeherKandukuri/Go_HotelReservationSys/internal/forms"

// Template Data holds any kind of data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	Data      map[string]interface{}
	CSRFToken string
	Success     string
	Warning   string
	Error     string
	Form      *forms.Form
	IsAuthentiated int
}
