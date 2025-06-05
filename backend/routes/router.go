package routes

import (
    "github.com/go-chi/chi/v5"
    "database/sql"
)

func RegisterAllRoutes(r chi.Router, db *sql.DB) {
    RegisterUserRoutes(r, db)
    RegisterTargetRoutes(r, db)
}