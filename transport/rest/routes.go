// Package rest implements the pageshot's REST API.
package rest

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/alyakimenko/pageshot/models"
	"github.com/sirupsen/logrus"
)

// ScreenshotService is a screenshot interface that is required for the Handler.
type ScreenshotService interface {
	Screenshot(ctx context.Context, opts models.ScreenshotOptions) (io.Reader, string, error)
}

// Handler is an REST HTTP handler.
type Handler struct {
	mux *http.ServeMux

	logger            *logrus.Logger
	screenshotService ScreenshotService
}

// HandlerParams is an incoming params for the NewHandler function.
type HandlerParams struct {
	Logger            *logrus.Logger
	ScreenshotService ScreenshotService
}

// NewHandler returns net http.Handler.
func NewHandler(params HandlerParams) *Handler {
	mux := http.NewServeMux()

	handler := &Handler{
		mux:               mux,
		logger:            params.Logger,
		screenshotService: params.ScreenshotService,
	}

	handler.initRoutes()

	return handler
}

// ServeHTTP satisfies the http.Handler interface.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// initRoutes initializes all routes for the Handler.
func (h *Handler) initRoutes() {
	h.mux.HandleFunc("/screenshot", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		h.screenshot(w, r)
	})

	start := time.Now()
	h.mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		h.health(start)(w, r)
	})
}
