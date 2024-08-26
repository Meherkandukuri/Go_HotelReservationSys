package models

import "time"

type User struct {
	ID        int
	FullName  string
	Email     string
	Password  string
	Role      int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Bungalow struct {
	ID           int
	BungalowName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Restriction struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservations is the model of a reservation
type Reservation struct {
	ID         int
	FullName   string
	Email      string
	Phone      string
	Password   string
	StartDate  time.Time
	EndDate    time.Time
	BungalowID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Bungalow   Bungalow
	Processed  int
}

// BungalowRestrictions
type BungalowRestriction struct {
	ID            int
	StartDate     time.Time
	EndDate       time.Time
	BungalowID    int
	ReservationID int
	RestrictionID int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Bungalow      Bungalow
	Reservation   Reservation
	Restriction   Restriction
}
