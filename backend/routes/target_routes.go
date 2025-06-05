package routes

import (
	"database/sql"
	"uptime-monitor/handler"
	"uptime-monitor/repository"
	"uptime-monitor/service"

	"github.com/go-chi/chi/v5"
)

func RegisterTargetRoutes(r chi.Router, db *sql.DB) {
	targetRepo := repository.NewTargetRepository(db)
	targetService := service.NewTargetService(targetRepo)
	targetHandler := handler.NewTargetHandler(targetService)

	r.Route("/targets", func(r chi.Router) {
		r.Get("/", targetHandler.GetAllTargets)
		r.Post("/", targetHandler.CreateTarget)
	})
}
