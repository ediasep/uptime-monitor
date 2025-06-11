package routes

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func RegisterAllRoutes(r chi.Router, db *sql.DB) {
	RegisterUserRoutes(r, db)
	RegisterTargetRoutes(r, db)
	RegisterTargetLogRoutes(r, db)
}
