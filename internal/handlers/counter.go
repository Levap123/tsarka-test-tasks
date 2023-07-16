package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *Handlers) getCounter(w http.ResponseWriter, r *http.Request) {
	counter, err := h.services.CounterServiceI.Get(r.Context())
	if err != nil {
		log.Printf("can not get counter: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("can not get counter during internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(counter)))
}

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

func (h *Handlers) addCounter(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(chi.URLParam(r, "num"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(errNotFound)
		return
	}

	if err := h.services.CounterServiceI.Add(r.Context(), num); err != nil {
		log.Printf("can not add to counter: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("can not add to counter during internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("counter changed"))
}
