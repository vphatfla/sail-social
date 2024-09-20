package router

import (
	"net/http"
	influencerhandler "vphatlfa/booster-hub/handler/influencerHandler"

	"github.com/go-chi/chi/v5"
)

func influencerRouter() http.Handler {
	r := chi.NewRouter()

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/", influencerhandler.DefaultHandler)
	})
	return r
}
