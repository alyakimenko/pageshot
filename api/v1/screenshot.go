package v1

import (
	"net/http"

	"github.com/alyakimenko/pageshot/models"
	"github.com/gin-gonic/gin"
)

// screenshotRequest is an HTTP request model for the GET /v1/screenshot route.
type screenshotRequest struct {
	URL      string `form:"url" binding:"required"`
	Width    int    `form:"width"`
	Height   int    `form:"height"`
	Quality  int    `form:"quality"`
	Fullpage bool   `form:"fullpage"`
}

// screenshot is an HTTP handler for GET /v1/screenshot route.
func (h *Handler) screenshot(c *gin.Context) {
	var req screenshotRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	image, contentType, err := h.ScreenshotService.Screenshot(c, models.ScreenshotOptions{
		URL:      req.URL,
		Width:    req.Width,
		Height:   req.Height,
		Quality:  req.Quality,
		Fullpage: req.Fullpage,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.Data(http.StatusOK, contentType, image)
}
