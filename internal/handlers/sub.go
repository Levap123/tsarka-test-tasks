package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *Handlers) subCounter(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(chi.URLParam(r, "num"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(errNotFound)
		return
	}

	if err := h.services.CounterServiceI.Sub(r.Context(), num); err != nil {
		log.Printf("can not substract counter: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("can not substract counter during internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("counter changed"))
}
