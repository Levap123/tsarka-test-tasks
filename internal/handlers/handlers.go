package handlers

import "github.com/Levap123/tsarka-test-tasks/internal/services"

type Handlers struct {
	services *services.Services
}

var (
	errReadBody = []byte("could not read request body")
	errNotFound = []byte("not found")
)

func NewHandlers(s *services.Services) *Handlers {
	return &Handlers{
		services: s,
	}
}
