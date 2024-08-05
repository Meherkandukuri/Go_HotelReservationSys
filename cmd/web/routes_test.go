package main

import (
	"testing"

	"github.com/MeherKandukuri/reservationSystem_Go/internal/config"

	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// all fine nothing to do
	default:
		t.Errorf("Type mismatch: Expected *chi.Mux, got %T", v)
	}

}
