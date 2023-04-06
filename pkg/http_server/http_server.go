package httpserver

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	http *http.Server
}

func (s *Server) Start(routes http.Handler) error {
	http := &http.Server{
		Handler:        routes,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	s.http = http

	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
