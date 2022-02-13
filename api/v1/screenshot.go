package v1

import (
	"net/http"

	"github.com/alyakimenko/pageshot/models"
	"github.com/gorilla/schema"
)

// screenshotRequest is an HTTP request model for the GET /v1/screenshot route.
type screenshotRequest struct {
	URL      string `schema:"url,required"`
	Width    int    `schema:"width"`
	Height   int    `schema:"height"`
	Quality  int    `schema:"quality"`
	Fullpage bool   `schema:"fullpage"`
}

// screenshot is an HTTP handler for GET /v1/screenshot route.
func (h *Handler) screenshot(w http.ResponseWriter, r *http.Request) {
	var req screenshotRequest
	if err := schema.NewDecoder().Decode(&req, r.URL.Query()); err != nil {
		h.logger.Errorf("failed to decode query: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	image, contentType, err := h.screenshotService.Screenshot(r.Context(), models.ScreenshotOptions{
		URL:      req.URL,
		Width:    req.Width,
		Height:   req.Height,
		Quality:  req.Quality,
		Fullpage: req.Fullpage,
	})
	if err != nil {
		h.logger.Errorf("failed to take a screenshot: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", contentType)
	_, err = w.Write(image)
	if err != nil {
		h.logger.Errorf("failed to write response: %s", err.Error())
	}
}
