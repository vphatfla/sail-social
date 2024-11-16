package router

import (
	"net/http"
	"vphatlfa/booster-hub/customMiddleware"
	"vphatlfa/booster-hub/handler/getSearchHandler"

	"github.com/go-chi/chi/v5"
)

func generalRouter() http.Handler {
	r := chi.NewRouter()

	// private routes
	r.Group(func(r chi.Router) {
		r.Use(customMiddleware.JWTMiddleware)
		r.Get("/post", getSearchHandler.GetPostHandler)
	})

	return r
}
