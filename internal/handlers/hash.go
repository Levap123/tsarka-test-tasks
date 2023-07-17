package handlers

import (
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *Handlers) startHashCalc(w http.ResponseWriter, r *http.Request) {
	hash, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("can not read request body"))
		return
	}

	ID, err := h.services.HashServiceI.GenerateHash(r.Context(), string(hash))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(ID)))
}

func (h *Handlers) getHash(w http.ResponseWriter, r *http.Request) {
	ID, err := strconv.Atoi(chi.URLParam(r, "hash_id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}

	hash, err := h.services.HashServiceI.Get(r.Context(), ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(hash))
}
