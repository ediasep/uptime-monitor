package routes

import (
	"database/sql"
	"uptime-monitor/handler"
	target "uptime-monitor/repository/target"
	"uptime-monitor/service"

	"github.com/go-chi/chi/v5"
)

func RegisterTargetRoutes(r chi.Router, db *sql.DB) {
	targetRepo := target.NewTargetRepository(db)
	targetService := service.NewTargetService(targetRepo)
	targetHandler := handler.NewTargetHandler(targetService)

	r.Route("/targets", func(r chi.Router) {
		r.Get("/", targetHandler.GetAllTargets)
		r.Post("/", targetHandler.CreateTarget)
	})
}
