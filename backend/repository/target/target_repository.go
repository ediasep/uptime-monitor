// repository/target_repository.go
package target

import (
	"database/sql"
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

func (r *targetRepository) List() ([]model.Target, error) {
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
