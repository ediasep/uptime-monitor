package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"uptime-monitor/repository/target"
	"uptime-monitor/repository/targetlog"
	"uptime-monitor/routes"
	"uptime-monitor/service"
	"uptime-monitor/storage"

	_ "uptime-monitor/docs" // swag docs

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Initialize the database
	db := storage.InitDatabase()
	defer db.Close()

	// Initialize repositories as interfaces
	var targetRepo target.TargetRepository = target.NewTargetRepository(db)
	var logRepo targetlog.TargetLogRepository = targetlog.NewTargetLogRepository(db)

	// Initialize services
	eventSvc := service.NewEventService()
	uptimeChecker := service.NewUptimeCheckerService(targetRepo, logRepo, eventSvc, time.Minute)
	uptimeChecker.Start()

	// Create router and register routes
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Uptime Monitor API!"))
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Route registration
	routes.RegisterAllRoutes(r, db)

	// Start server
	fmt.Println("\033[1m - API Server running on http://localhost:8080 \033[0m")
	log.Fatal(http.ListenAndServe(":8080", r))
}
