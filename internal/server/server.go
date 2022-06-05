package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	addr       string
	handler    http.Handler
	httpServer *http.Server
}

func NewServer(addr string, handler http.Handler) *Server {
	return &Server{
		addr:    addr,
		handler: handler,
	}
}

func (s *Server) Run() error {
	s.httpServer = &http.Server{
		Addr:           s.addr,
		Handler:        s.handler,
		MaxHeaderBytes: 1 << 20, // MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
