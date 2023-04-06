package handlers

import (
	"io"
	"net/http"
)

func (h *Handlers) findSubstring(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(errMethodNotAllowed)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errReadBody)
		return
	}

	subString := h.services.AlgosServiceI.FindSubstring(string(bytes))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(subString))
}
