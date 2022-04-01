package rest

import (
	"encoding/json"
	"net/http"
	"time"
)

// uptimeResponse is a response model for the GET /health route.
type uptimeResponse struct {
	Uptime string `json:"uptime"`
}

// health is an HTTP handler for GET /health route.
func (h *Handler) health(start time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp := uptimeResponse{
			Uptime: time.Since(start).String(),
		}

		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			h.logger.Errorf("encode uptime: %s", err.Error())
		}
	}
}
