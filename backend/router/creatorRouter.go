package router

import (
	"net/http"
	"vphatlfa/booster-hub/handler/creatorHandler"

	"github.com/go-chi/chi/v5"
)

func CreatorRouter() http.Handler {
	r := chi.NewRouter()

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/", creatorHandler.DefaultHandler)
	})
	return r
}
