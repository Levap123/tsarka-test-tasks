package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type createUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (h *Handlers) createUser(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respBytes := marshalJSON(map[string]string{"msg": "can not unmarshal response body"})

		sendJSON(w, http.StatusBadRequest, respBytes)
		return
	}

	ID, err := h.services.UserServiceI.Create(r.Context(), req.FirstName, req.LastName)
	if err != nil {
		log.Printf("can not create user: %v", err)
		respBytes := marshalJSON(map[string]string{"msg": "internal error"})

		sendJSON(w, http.StatusInternalServerError, respBytes)
		return
	}

	respBytes := marshalJSON(map[string]int{"ID": ID})
	sendJSON(w, http.StatusOK, respBytes)
}

func (h *Handlers) getUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		respBytes := marshalJSON(map[string]string{"msg": "not found"})
		sendJSON(w, http.StatusNotFound, respBytes)
		return
	}

	user, err := h.services.UserServiceI.Get(r.Context(), userID)
	if err != nil {
		log.Printf("can not get user: %v", err)
		respBytes := marshalJSON(map[string]string{"msg": "internal error"})

		sendJSON(w, http.StatusInternalServerError, respBytes)
		return
	}

	respBytes := marshalJSON(user)

	sendJSON(w, http.StatusOK, respBytes)
}

func (h *Handlers) deleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		respBytes := marshalJSON(map[string]string{"msg": "not found"})
		sendJSON(w, http.StatusNotFound, respBytes)
		return
	}

	if err := h.services.UserServiceI.Delete(r.Context(), userID); err != nil {
		log.Printf("can not delete user: %v", err)
		respBytes := marshalJSON(map[string]string{"msg": "internal error"})

		sendJSON(w, http.StatusInternalServerError, respBytes)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}

type updateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (h *Handlers) updateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		respBytes := marshalJSON(map[string]string{"msg": "not found"})
		sendJSON(w, http.StatusNotFound, respBytes)
		return
	}

	var req updateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respBytes := marshalJSON(map[string]string{"msg": "can not unmarshal response body"})

		sendJSON(w, http.StatusBadRequest, respBytes)
		return
	}

	if err := h.services.UserServiceI.Update(r.Context(), userID, req.FirstName, req.LastName); err != nil {
		log.Printf("can not update user: %v", err)
		respBytes := marshalJSON(map[string]string{"msg": "internal error"})

		sendJSON(w, http.StatusInternalServerError, respBytes)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}
