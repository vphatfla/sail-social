package router

import (
	"net/http"
	"vphatlfa/booster-hub/handler"
	"vphatlfa/booster-hub/router/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

func MainRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger.Logger()))
	r.Mount("/api", redirectRouter())
	return r
}

func redirectRouter() http.Handler {
	r := chi.NewRouter()

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/helloworld", handler.HelloWorldHandler)
	})

	r.Mount("/Creator", CreatorRouter())
	r.Mount("/business", businessRouter())
	r.Mount("/admin", adminRouter())
	return r
}
