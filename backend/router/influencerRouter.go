package router

import (
	"net/http"
	"vphatlfa/booster-hub/router/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

func influencerRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger.Logger()))
	return r
}
