package helper

import (
	"fmt"
	"os"
)

type Config struct {
	Driver string
	DSN    string
}

func LoadConfig() (*Config, error) {
	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		return nil, fmt.Errorf("DB_DRIVER environment variable is required")
	}
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return nil, fmt.Errorf("DB_DSN environment variable is required")
	}

	return &Config{
		Driver: driver,
		DSN:    dsn,
	}, nil
}
