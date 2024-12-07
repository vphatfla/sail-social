package router

import (
	"net/http"
	"vphatlfa/booster-hub/customMiddleware"
	businesshandler "vphatlfa/booster-hub/handler/businessHandler"
	"vphatlfa/booster-hub/handler/postHandler"

	"github.com/go-chi/chi/v5"
)

func businessRouter() http.Handler {
	r := chi.NewRouter()

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/", businesshandler.DefaultHandler)
		r.Post("/sign-up", businesshandler.SignUpHandler)
		r.Post("/log-in", businesshandler.LoginHandler)
	})

	// Private routes
	r.Group(func(r chi.Router) {
		r.Use(customMiddleware.JWTMiddleware)
		r.Get("/test", businesshandler.TestHandler)
		r.Post("/onboarding", businesshandler.OnboadingHandler)
		r.Post("/post-new", postHandler.CreateNewPostHandler)
		r.Post("/post-update", postHandler.UpdatePostHandler)
		r.Post("/search-business-cred", businesshandler.SearchBusinessCredHandler)
		r.Post("/search-business-info", businesshandler.SearchBusinessInfoHandler)
	})
	return r
}
