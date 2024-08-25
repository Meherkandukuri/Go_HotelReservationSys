package dbrepo

import (
	"database/sql"

	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/config"
	"github.com/MeherKandukuri/Go_HotelReservationSys/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostGresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
