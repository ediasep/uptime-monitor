package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

// UptimeResponse represents the response for uptime checks
type UptimeResponse struct {
	URL        string `json:"url"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Duration   int64  `json:"duration"`
}

// CheckUptimeHandler godoc
// @Summary      Check website uptime
// @Description  Checks the uptime of a given website URL
// @Tags         uptime
// @Accept       json
// @Produce      json
// @Param        url  query     string  true  "Website URL to check"
// @Success      200  {object}  handler.UptimeResponse
// @Failure      400  {string}  string "Missing url parameter"
// @Router       /uptime [get]
func CheckUptimeHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing url parameter", http.StatusBadRequest)
		return
	}

	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	result := UptimeResponse{
		URL:      url,
		Duration: duration.Milliseconds(),
	}

	if err != nil {
		result.Status = "down"
		result.StatusCode = 0
	} else {
		defer resp.Body.Close()
		result.Status = "up"
		result.StatusCode = resp.StatusCode
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
