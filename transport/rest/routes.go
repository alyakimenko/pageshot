// Package rest implements the pageshot's REST API.
package rest

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/alyakimenko/pageshot/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// ScreenshotService is a screenshot interface that is required for the Handler.
type ScreenshotService interface {
	Screenshot(ctx context.Context, opts models.ScreenshotOptions) (io.Reader, string, error)
}

// Handler is an REST HTTP handler.
type Handler struct {
	router *mux.Router

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
	router := mux.NewRouter()

	handler := &Handler{
		router:            router,
		logger:            params.Logger,
		screenshotService: params.ScreenshotService,
	}

	handler.initRoutes()

	router.Use(handler.recovery)

	return handler
}

// ServeHTTP satisfies the http.Handler interface.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

// initRoutes initializes all routes for the Handler.
func (h *Handler) initRoutes() {
	h.router.HandleFunc("/screenshot", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		h.screenshot(w, r)
	})

	start := time.Now()
	h.router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		h.health(start)(w, r)
	})
}
