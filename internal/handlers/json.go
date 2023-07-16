package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func marshalJSON(v any) []byte {
	respBytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("can not marshal json: %v\n", err)
		return nil
	}

	return respBytes
}

func sendJSON(w http.ResponseWriter, status int, respBytes []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write(respBytes)
}
