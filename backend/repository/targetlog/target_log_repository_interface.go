// repository/target_repository_interface.go
package targetlog

import "uptime-monitor/model"

type TargetLogRepository interface {
	Save(log model.TargetLog) error
	CountRecentFailures(targetID string, limit int) (int, error)
	ListByTargetID(targetID string) ([]model.TargetLog, error)
	DeleteByTargetID(targetID string) error
	GetDailyUptimePercentageByTargetID(targetID string) ([]model.DailyUptimeResponse, error)
}
