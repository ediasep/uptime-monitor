package helper

import (
	"os"
)

type Config struct {
	Driver          string
	DSN             string
	TimestampLayout string
}

func LoadConfig() (*Config, error) {
	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = "sqlite3" // Default to sqlite3 if not set
	}
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "./targets.db" // Default to ./targets.db if not set
	}
	layout := os.Getenv("TIMESTAMP_LAYOUT")
	if layout == "" {
		layout = "2006-01-02 15:04:05.999999-07:00" // Default to Go's layout for datetime with timezone
	}

	return &Config{
		Driver:          driver,
		DSN:             dsn,
		TimestampLayout: layout,
	}, nil
}
