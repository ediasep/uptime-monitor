// repository/target_repository.go
package target

import (
	"database/sql"
	"time"
	"uptime-monitor/helper"
	"uptime-monitor/model"
)

type targetRepository struct {
	db *sql.DB
}

func NewTargetRepository(db *sql.DB) TargetRepository {
	return &targetRepository{db: db}
}

func (r *targetRepository) Add(name, url string, interval int) (model.Target, error) {
	target := model.NewTarget(name, url, interval)

	_, err := r.db.Exec(`
		INSERT INTO targets (id, name, url, interval, created_at)
		VALUES (?, ?, ?, ?, ?)`,
		target.ID, target.Name, target.URL, target.Interval, target.CreatedAt)

	return target, err
}

func (r *targetRepository) Update(id, name, url string, interval int) (model.Target, error) {
	_, err := r.db.Exec(`
		UPDATE targets
		SET name = ?, url = ?, interval = ?
		WHERE id = ?`,
		name, url, interval, id)

	// Optionally, fetch the updated target to return
	var target model.Target
	errScan := r.db.QueryRow(`SELECT id, name, url, interval, created_at FROM targets WHERE id = ?`, id).
		Scan(&target.ID, &target.Name, &target.URL, &target.Interval, &target.CreatedAt)
	if errScan != nil {
		return target, errScan
	}

	return target, err
}

func (r *targetRepository) List() ([]model.TargetListDto, error) {
	// Load config to get the timestamp layout
	config, _ := helper.LoadConfig()
	layout := config.TimestampLayout

	rows, err := r.db.Query(`
		SELECT 
			t.id, t.name, t.url, t.interval, t.created_at,
			l.timestamp as last_checked_at,
			l.status
		FROM targets t
		LEFT JOIN (
			SELECT target_id, status, MAX(timestamp) as timestamp
			FROM target_logs
			GROUP BY target_id
		) l ON t.id = l.target_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var targets []model.TargetListDto
	for rows.Next() {
		var t model.TargetListDto
		var lastCheckedAt sql.NullString
		var status sql.NullString

		err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.URL,
			&t.Interval,
			&t.CreatedAt,
			&lastCheckedAt,
			&status,
		)
		if err != nil {
			return nil, err
		}

		// Convert lastCheckedAt to *time.Time using layout from config
		if lastCheckedAt.Valid && lastCheckedAt.String != "" {
			parsed, err := time.Parse(layout, lastCheckedAt.String)
			if err == nil {
				t.LastCheckedAt = &parsed
			}
		}

		// Convert status to *string
		if status.Valid {
			t.LastStatus = &status.String
		}

		targets = append(targets, t)
	}
	return targets, nil
}

func (r *targetRepository) GetByID(id string) (model.Target, error) {
	var target model.Target
	err := r.db.QueryRow(`SELECT id, name, url, interval, created_at FROM targets WHERE id = ?`, id).
		Scan(&target.ID, &target.Name, &target.URL, &target.Interval, &target.CreatedAt)
	return target, err
}

func (r *targetRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM targets WHERE id = ?`, id)
	return err
}
