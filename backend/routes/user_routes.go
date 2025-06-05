package routes

import (
	"database/sql"
	"uptime-monitor/handler"
	"uptime-monitor/repository"
	"uptime-monitor/service"

	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutes(r chi.Router, db *sql.DB) {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetUser)
		r.Post("/", userHandler.CreateUser)
	})
}
