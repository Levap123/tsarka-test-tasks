package httpserver

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	http *http.Server
}

// hardcoded because i too lazy for adding configs
func New(routes http.Handler, addr string) *Server {
	return &Server{
		http: &http.Server{
			Addr:           addr,
			Handler:        routes,
			ReadTimeout:    time.Second * 10,
			WriteTimeout:   time.Second * 10,
			MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		},
	}
}

func (s *Server) Start() error {
	return s.http.ListenAndServe()
}

func (srv *Server) SetKeepAlivesEnabled(v bool) {
	srv.http.SetKeepAlivesEnabled(v)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
