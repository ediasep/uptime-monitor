package repository

import (
	"database/sql"
	"uptime-monitor/model"
)

type TargetLogRepository struct {
	db *sql.DB
}

func NewTargetLogRepository(db *sql.DB) *TargetLogRepository {
	return &TargetLogRepository{db: db}
}

func (r *TargetLogRepository) Save(log model.TargetLog) error {
	_, err := r.db.Exec(
		"INSERT INTO target_logs (id, target_id, status, timestamp) VALUES (?, ?, ?, ?)",
		log.ID, log.TargetID, log.Status, log.Timestamp,
	)
	return err
}

func (r *TargetLogRepository) CountRecentFailures(targetID string, limit int) (int, error) {
	rows, err := r.db.Query(
		`SELECT status FROM target_logs 
		 WHERE target_id = ? 
		 ORDER BY timestamp DESC 
		 LIMIT ?`, targetID, limit,
	)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	failures := 0
	for rows.Next() {
		var status string
		if err := rows.Scan(&status); err != nil {
			return 0, err
		}
		if status != "DOWN" {
			break
		}
		failures++
	}

	return failures, nil
}
