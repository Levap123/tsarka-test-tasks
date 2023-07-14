package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (h *Handlers) InitRoutes() http.Handler {
	router := chi.NewRouter()

	router.Post("/rest/substr/find", h.findSubstring)
	router.Post("/rest/email/check", h.checkStringOnEmailPattern)

	router.Post("/rest/counter/add/{num}", h.addCounter)
	router.Post("/rest/counter/sub/{num}", h.subCounter)
	router.Get("/rest/counter/val", h.getCounter)

	router.Get("/swagger/*", httpSwagger.Handler())

	return router
}
