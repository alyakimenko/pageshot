// Package v1 implements the v1 of the pageshot's REST API.
package v1

import (
	"context"
	"net/http"

	"github.com/alyakimenko/pageshot/models"
	"github.com/sirupsen/logrus"
)

// ScreenshotService is a screenshot interface that is required for the Handler.
type ScreenshotService interface {
	Screenshot(ctx context.Context, opts models.ScreenshotOptions) ([]byte, string, error)
}

// Handler is an HTTP handler for v1 routes.
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
	h.mux.HandleFunc("/v1/screenshot", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		h.screenshot(w, r)
	})
}
