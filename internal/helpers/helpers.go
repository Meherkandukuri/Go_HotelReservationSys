package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/config"
)

var app *config.AppConfig

// NewHelpers make application wide settings available in package helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error! status:", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}


// IsAuthenticated figures determines if an authenticated user exists in the session data

func IsAuthenticated(r *http.Request) bool {
	exists := app.Session.Exists(r.Context(), "user_id")
	return exists
}