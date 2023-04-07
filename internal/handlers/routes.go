package handlers

import (
	"net/http"
)

func (h *Handlers) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/rest/substr/find", h.findSubstring)
	mux.HandleFunc("/rest/email/check", h.checkStringOnEmailPattern)
	return mux
}
