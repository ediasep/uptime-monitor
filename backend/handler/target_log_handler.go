package handler

import (
	"encoding/json"
	"net/http"

	"uptime-monitor/service"

	"github.com/go-chi/chi/v5"
)

type TargetLogHandler struct {
	service *service.TargetLogService
}

func NewTargetLogHandler(service *service.TargetLogService) *TargetLogHandler {
	return &TargetLogHandler{service: service}
}

// GetLogsByTargetID godoc
// @Summary      Get logs by target ID
// @Description  Retrieve all logs for a specific target
// @Tags         target-logs
// @Accept       json
// @Produce      json
// @Param        id   path      string   true  "Target ID"
// @Success      200  {array}  model.TargetLog
// @Failure      404  {string}  string "Logs not found"
// @Failure      500  {string}  string "Failed to fetch logs"
// @Router       /targets/{id}/logs [get]
func (h *TargetLogHandler) GetLogsByTargetID(w http.ResponseWriter, r *http.Request) {
	targetID := chi.URLParam(r, "id")
	logs, err := h.service.GetLogsByTargetID(targetID)
	if err != nil {
		http.Error(w, "Failed to fetch logs", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

// DeleteLogsByTargetID godoc
// @Summary      Delete logs by target ID
// @Description  Delete all logs for a specific target
// @Tags         target-logs
// @Accept       json
// @Produce      json
// @Param        id   path      string   true  "Target ID"
// @Success      204  {string}  string "Logs deleted"
// @Failure      500  {string}  string "Failed to delete logs"
// @Router       /targets/{id}/logs [delete]
func (h *TargetLogHandler) DeleteLogsByTargetID(w http.ResponseWriter, r *http.Request) {
	targetID := chi.URLParam(r, "id")
	err := h.service.DeleteLogsByTargetID(targetID)
	if err != nil {
		http.Error(w, "Failed to delete logs", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
