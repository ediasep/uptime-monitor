package handler

import (
	"encoding/json"
	"net/http"
	"uptime-monitor/service"
)

// TargetHandler handles target-related requests
type TargetHandler struct {
	service *service.TargetService
}

func NewTargetHandler(svc *service.TargetService) *TargetHandler {
	return &TargetHandler{service: svc}
}

type CreateTargetRequest struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Interval int    `json:"interval"`
}

// CreateTarget godoc
// @Summary      Create a new target
// @Description  Create a new uptime monitoring target
// @Tags         targets
// @Accept       json
// @Produce      json
// @Param        target  body      handler.CreateTargetRequest  true  "Target payload"
// @Success      200     {object}  model.Target
// @Failure      400     {string}  string "Missing or invalid fields"
// @Failure      500     {string}  string "Failed to save target"
// @Router       /targets [post]
func (h *TargetHandler) CreateTarget(w http.ResponseWriter, r *http.Request) {
	var req CreateTargetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.URL == "" || req.Interval < 5 {
		http.Error(w, "Missing or invalid fields", http.StatusBadRequest)
		return
	}

	target, err := h.service.CreateTarget(req.Name, req.URL, req.Interval)
	if err != nil {
		http.Error(w, "Failed to save target", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(target)
}

// GetAllTargets godoc
// @Summary      Get all targets
// @Description  Retrieve all uptime monitoring targets
// @Tags         targets
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Target
// @Failure      500  {string}  string "Failed to fetch targets"
// @Router       /targets [get]
func (h *TargetHandler) GetAllTargets(w http.ResponseWriter, r *http.Request) {
	targets, err := h.service.GetAllTargets()
	if err != nil {
		http.Error(w, "Failed to fetch targets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(targets)
}
