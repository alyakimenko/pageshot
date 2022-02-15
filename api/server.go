package api

import (
	"context"
	"net/http"
	"time"

	"github.com/alyakimenko/pageshot/config"
)

const (
	// defaultServerShutdownTimeout is a default timeout to wait before shutting down the server.
	defaultServerShutdownTimeout = 5 * time.Second
)

// Server is main HTTP server.
type Server struct {
	server *http.Server
}

// ServerParams is an incoming params for the NewServer function.
type ServerParams struct {
	Config  config.ServerConfig
	Handler http.Handler
}

// NewServer creates new instance of the Server.
func NewServer(params ServerParams) *Server {
	return &Server{
		server: &http.Server{
			Addr:         params.Config.Addr(),
			Handler:      params.Handler,
			ReadTimeout:  params.Config.ReadTimeout,
			WriteTimeout: params.Config.WriteTimeout,
			IdleTimeout:  params.Config.IdleTimeout,
		},
	}
}

// Start starts the server.
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

// Shutdown gracefully shut downs the server.
func (s *Server) Shutdown(ctx context.Context) error {
	shutdownCtx, cancel := context.WithTimeout(ctx, defaultServerShutdownTimeout)
	defer cancel()

	s.server.SetKeepAlivesEnabled(false)

	return s.server.Shutdown(shutdownCtx)
}
