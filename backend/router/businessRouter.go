package router

import (
	"net/http"
	businesshandler "vphatlfa/booster-hub/handler/businessHandler"

	"github.com/go-chi/chi/v5"
)

func businessRouter() http.Handler {
	r := chi.NewRouter()

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/", businesshandler.DefaultHandler)
	})
	return r
}
