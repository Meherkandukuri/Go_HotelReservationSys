package repository

import (
	"time"

	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertBungalowRestriction(r models.BungalowRestriction) error
	SearchAvailibilityByDatesByBungalowID(start, end time.Time, bungalowID int) (bool, error)
	SearchAvailabilityByDatesForAllBungalows(start, end time.Time) ([]models.Bungalow, error)
	GetBungalowByID(id int) (models.Bungalow, error)
}
