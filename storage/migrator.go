package storage

import (
	"database/sql"
	"fmt"
	"os"
)

func RunInitSchema(db *sql.DB, path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}
	_, err = db.Exec(string(content))
	if err != nil {
		return fmt.Errorf("failed to run schema: %w", err)
	}
	fmt.Println("Database schema initialized successfully")
	return nil
}
