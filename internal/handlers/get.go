package handlers

import (
	"log"
	"net/http"
	"strconv"
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
