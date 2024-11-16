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
		r.Get("/post-creator-applied", getSearchHandler.GetPostAppliedByCreatorHandler)
		r.Get("/post-creator-saved", getSearchHandler.GetPostSavedByCreatorHandler)
		r.Get("/post-business", getSearchHandler.GetPostByBusinessHandler)
		// search-get creator
		r.Get("/creator-search", getSearchHandler.GetCreatorHandler)
		// search-get business
		r.Get("/business-search", getSearchHandler.GetBusinessHandler)
	})

	return r
}
