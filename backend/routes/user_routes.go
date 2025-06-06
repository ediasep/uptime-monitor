package routes

import (
	"database/sql"
	"uptime-monitor/handler"
	repository "uptime-monitor/repository/user"
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
