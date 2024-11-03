package router

import (
	"net/http"
	"vphatlfa/booster-hub/handler"
	"vphatlfa/booster-hub/router/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		// AllowedOrigins: []string{"https://*", "http://*"},
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/helloworld", handler.HelloWorldHandler)
	})

	r.Mount("/creator", creatorRouter())
	r.Mount("/business", businessRouter())
	r.Mount("/admin", adminRouter())
	r.Mount("/general", generalRouter())
	return r
}
