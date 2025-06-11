package routes

import (
	"database/sql"
	"uptime-monitor/handler"
	"uptime-monitor/repository/targetlog"
	"uptime-monitor/service"

	"github.com/go-chi/chi/v5"
)

func RegisterTargetLogRoutes(r chi.Router, db *sql.DB) {
	targetLogRepo := targetlog.NewTargetLogRepository(db)
	targetLogService := service.NewTargetLogService(targetLogRepo)
	targetLogHandler := handler.NewTargetLogHandler(targetLogService)

	r.Route("/targets/{id}/logs", func(r chi.Router) {
		r.Get("/", targetLogHandler.GetLogsByTargetID)
		r.Delete("/", targetLogHandler.DeleteLogsByTargetID)
	})
}
