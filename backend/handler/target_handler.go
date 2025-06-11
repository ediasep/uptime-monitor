package handler

import (
	"encoding/json"
	"net/http"
	"uptime-monitor/service"

	"github.com/go-chi/chi/v5"
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

// UpdateTarget godoc
// @Summary      Update a target
// @Description  Update an existing uptime monitoring target
// @Tags         targets
// @Accept       json
// @Produce      json
// @Param        id      path      string                    true  "Target ID"
// @Param        target  body      handler.CreateTargetRequest  true  "Target payload"
// @Success      200     {object}  model.Target
// @Failure      400     {string}  string "Missing or invalid fields"
// @Failure      404     {string}  string "Target not found"
// @Failure      500     {string}  string "Failed to update target"
// @Router       /targets/{id} [put]
func (h *TargetHandler) UpdateTarget(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req CreateTargetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.URL == "" || req.Interval < 5 {
		http.Error(w, "Missing or invalid fields", http.StatusBadRequest)
		return
	}

	target, err := h.service.UpdateTarget(id, req.Name, req.URL, req.Interval)
	if err != nil {
		if err == service.ErrTargetNotFound {
			http.Error(w, "Target not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to update target", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(target)
}

// GetTargetByID godoc
// @Summary      Get a target by ID
// @Description  Retrieve a single uptime monitoring target by its ID
// @Tags         targets
// @Accept       json
// @Produce      json
// @Param        id   path      string         true  "Target ID"
// @Success      200  {object}  model.Target
// @Failure      404  {string}  string "Target not found"
// @Failure      500  {string}  string "Failed to fetch target"
// @Router       /targets/{id} [get]
func (h *TargetHandler) GetTargetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	target, err := h.service.GetTargetByID(id)
	if err != nil {
		if err == service.ErrTargetNotFound {
			http.Error(w, "Target not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to fetch target", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(target)
}

// DeleteTarget godoc
// @Summary      Delete a target
// @Description  Delete an uptime monitoring target by its ID
// @Tags         targets
// @Accept       json
// @Produce      json
// @Param        id   path      string   true  "Target ID"
// @Success      204  {string}  string "Target deleted"
// @Failure      404  {string}  string "Target not found"
// @Failure      500  {string}  string "Failed to delete target"
// @Router       /targets/{id} [delete]
func (h *TargetHandler) DeleteTarget(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.service.DeleteTarget(id)
	if err != nil {
		if err == service.ErrTargetNotFound {
			http.Error(w, "Target not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete target", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
