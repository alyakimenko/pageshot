package api

import (
	"context"
	"net/http"

	"github.com/alyakimenko/pageshot/config"
)

// Server is main HTTP server.
type Server struct {
	server *http.Server
}

// Params is an incoming params for the NewServer function.
type Params struct {
	Config  config.ServerConfig
	Handler http.Handler
}

// NewServer creates new instance of the Server.
func NewServer(params Params) *Server {
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
	return s.server.Shutdown(ctx)
}
