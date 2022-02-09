package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ScreenshotService is a screenshot interface that is required for the Handler.
type ScreenshotService interface {
	Screenshot(ctx context.Context, url string) ([]byte, error)
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

	group := handler.engine.Group("v1")
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

// screenshot is an HTTP handler for GET /screenshot route.
func (h *Handler) screenshot(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.Status(http.StatusBadRequest)

		return
	}

	image, err := h.ScreenshotService.Screenshot(c, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.Data(http.StatusOK, "image/png", image)
}
