package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/markledger/bookings/internal/config"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
	default:
		t.Error(fmt.Printf("Type is not chi.Mux it is:  %T", v))
	}
}
