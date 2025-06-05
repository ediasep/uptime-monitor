package storage

import (
	"database/sql"
	"log"
	"uptime-monitor/helper"

	"github.com/joho/godotenv"
)

func InitDatabase() *sql.DB {
	// Load .env file
	_ = godotenv.Load()

	config, err := helper.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	db, err := NewConnection(DBConfig{
		Driver: config.Driver,
		DSN:    config.DSN,
	})
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	if err := RunInitSchema(db, "migrations/init_schema.sql"); err != nil {
		log.Fatal("Schema init error:", err)
	}

	return db
}
