package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func adminRouter() http.Handler {
	r := chi.NewRouter()

	return r
}
