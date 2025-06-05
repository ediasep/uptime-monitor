package repository

import (
	"database/sql"
	"uptime-monitor/model"
)

type TargetRepository struct {
	db *sql.DB
}

func NewTargetRepository(db *sql.DB) *TargetRepository {
	return &TargetRepository{db: db}
}

func (r *TargetRepository) Add(name, url string, interval int) (model.Target, error) {
	target := model.NewTarget(name, url, interval)

	_, err := r.db.Exec(`
		INSERT INTO targets (id, name, url, interval, created_at)
		VALUES (?, ?, ?, ?, ?)`,
		target.ID, target.Name, target.URL, target.Interval, target.CreatedAt)

	return target, err
}

func (r *TargetRepository) List() ([]model.Target, error) {
	rows, err := r.db.Query(`SELECT id, name, url, interval, created_at FROM targets`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var targets []model.Target
	for rows.Next() {
		var t model.Target
		if err := rows.Scan(&t.ID, &t.Name, &t.URL, &t.Interval, &t.CreatedAt); err != nil {
			return nil, err
		}
		targets = append(targets, t)
	}
	return targets, nil
}
