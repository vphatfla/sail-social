package router

import (
	"net/http"
	"vphatlfa/booster-hub/customMiddleware"
	"vphatlfa/booster-hub/handler/creatorHandler"
	"vphatlfa/booster-hub/handler/postHandler"

	"github.com/go-chi/chi/v5"
)

func creatorRouter() http.Handler {
	r := chi.NewRouter()

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/", creatorHandler.DefaultHandler)
		r.Post("/sign-up", creatorHandler.SignUpHandler)
		r.Post("/log-in", creatorHandler.LoginHandler)
	})

	// Private Routes
	r.Group(func(r chi.Router) {
		r.Use(customMiddleware.JWTMiddleware)
		r.Get("/test", creatorHandler.TestHandler)
		r.Post("/onboarding", creatorHandler.OnboadingHandler)
		r.Post("/portfolio-new", creatorHandler.PostNewPortfolioHandler)
		r.Post("/post-apply", postHandler.NewApplyToPostHandler)
		r.Post("/post-update", postHandler.UpdateApplyToPostHandler)
	})
	return r
}
