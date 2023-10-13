package main

import (
	"fmt"
	"github.com/AlejandroDelg/webgo/internal/config"
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app *config.AppConfig

	mux := routes(app)

	switch v := mux.(type) {
	case *chi.Mux:
	default:
		t.Error(fmt.Println("type is not chi.mux, type is ", v))
	}
}
