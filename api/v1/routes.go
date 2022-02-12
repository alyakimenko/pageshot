// Package v1 implements the v1 of the pageshot's REST API.
package v1

import (
	"context"
	"net/http"

	"github.com/alyakimenko/pageshot/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	// v1PathPrefix is a path prefix that is used for v1 routes.
	v1PathPrefix = "v1"
)

// ScreenshotService is a screenshot interface that is required for the Handler.
type ScreenshotService interface {
	Screenshot(ctx context.Context, opts models.ScreenshotOptions) ([]byte, string, error)
}

// Handler is an HTTP handler for v1 routes.
type Handler struct {
	engine *gin.Engine

	Logger            *logrus.Logger
	ScreenshotService ScreenshotService
}

// HandlerParams is an incoming params for the NewHandler function.
type HandlerParams struct {
	Logger            *logrus.Logger
	ScreenshotService ScreenshotService
}

// NewHandler returns net http.Handler.
func NewHandler(params HandlerParams) *Handler {
	engine := gin.New()

	engine.Use(
		gin.Recovery(),
	)

	handler := &Handler{
		engine: engine,

		Logger:            params.Logger,
		ScreenshotService: params.ScreenshotService,
	}

	group := handler.engine.Group(v1PathPrefix)
	handler.initRoutes(group)

	return handler
}

// ServeHTTP satisfies the http.Handler interface.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.engine.ServeHTTP(w, r)
}

// initRoutes initializes all routes for the Handler.
func (h *Handler) initRoutes(group *gin.RouterGroup) {
	group.GET("/screenshot", h.screenshot)
}
