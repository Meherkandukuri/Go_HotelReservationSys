package config

import (
	"html/template"
	"log"

	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/models"
	"github.com/alexedwards/scs/v2"
)

// AppConfig is a struct holding die application's configuration

type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
