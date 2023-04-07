package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func (h *Handlers) checkStringOnEmailPattern(w http.ResponseWriter, r *http.Request) {
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

	emails := h.services.AlgosServiceI.EmailCheck(string(bytes))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", emails)))
}
