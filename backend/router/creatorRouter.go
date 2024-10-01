package router

import (
	"net/http"
	"vphatlfa/booster-hub/customMiddleware"
	"vphatlfa/booster-hub/handler/creatorHandler"

	"github.com/go-chi/chi/v5"
)

func creatorRouter() http.Handler {
	r := chi.NewRouter()

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/", creatorHandler.DefaultHandler)
		r.Post("/sign-up", creatorHandler.CreatorSignUpHandler)
		r.Post("/log-in", creatorHandler.LoginHandler)

	})

	// Private Routes
	r.Group(func(r chi.Router) {
		r.Use(customMiddleware.CreatorJWTMiddleware)
		r.Get("/test", creatorHandler.TestHandler)
		r.Post("/onboarding", creatorHandler.OnboadingHandler)
	})
	return r
}
