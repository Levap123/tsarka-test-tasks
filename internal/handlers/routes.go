package handlers

import (
	"net/http"
)

func (h *Handlers) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/rest/substr/find", h.findSubstring)

	return mux
}
