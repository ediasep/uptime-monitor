package targetlog

import (
	"database/sql"
	"time"
	"uptime-monitor/helper"
	"uptime-monitor/model"
)

type targetLogRepository struct {
	db *sql.DB
}

func NewTargetLogRepository(db *sql.DB) TargetLogRepository {
	return &targetLogRepository{db: db}
}

func (r *targetLogRepository) Save(log model.TargetLog) error {
	_, err := r.db.Exec(
		"INSERT INTO target_logs (id, target_id, status, timestamp) VALUES (?, ?, ?, ?)",
		log.ID, log.TargetID, log.Status, log.Timestamp,
	)
	return err
}

func (r *targetLogRepository) CountRecentFailures(targetID string, limit int) (int, error) {
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

func (r *targetLogRepository) ListByTargetID(targetID string) ([]model.TargetLog, error) {
	rows, err := r.db.Query(
		`SELECT id, target_id, status, timestamp FROM target_logs WHERE target_id = ? ORDER BY timestamp DESC`,
		targetID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []model.TargetLog
	for rows.Next() {
		var log model.TargetLog
		if err := rows.Scan(&log.ID, &log.TargetID, &log.Status, &log.Timestamp); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

func (r *targetLogRepository) DeleteByTargetID(targetID string) error {
	_, err := r.db.Exec(`DELETE FROM target_logs WHERE target_id = ?`, targetID)
	return err
}

func (r *targetLogRepository) GetDailyUptimePercentageByTargetID(targetID string) ([]model.DailyUptimeResponse, error) {
	rows, err := r.db.Query(`
        SELECT 
		    target_id,
            DATE(timestamp) AS date,
            SUM(CASE WHEN status = 'UP' THEN 1 ELSE 0 END) * 100.0 / COUNT(*) AS uptime_percentage
        FROM target_logs
        WHERE target_id = ?
            AND timestamp >= DATE('now', '-7 days')
        GROUP BY DATE(timestamp)
        ORDER BY date ASC
    `, targetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	config, _ := helper.LoadConfig()
	layout := config.DateLayout

	var results []model.DailyUptimeResponse
	for rows.Next() {
		var resp model.DailyUptimeResponse
		var dateStr string
		if err := rows.Scan(&resp.TargetID, &dateStr, &resp.UptimePercentage); err != nil {
			return nil, err
		}
		parsedDate, err := time.Parse(layout, dateStr)
		if err != nil {
			return nil, err
		}
		resp.Date = parsedDate
		results = append(results, resp)
	}
	return results, nil
}
