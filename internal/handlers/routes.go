package handlers

import (
	"net/http"

	_ "github.com/Levap123/tsarka-test-tasks/docs"

	"github.com/swaggo/http-swagger"
)

func (h *Handlers) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/rest/substr/find", h.findSubstring)
	mux.HandleFunc("/rest/email/check", h.checkStringOnEmailPattern)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	return mux
}
