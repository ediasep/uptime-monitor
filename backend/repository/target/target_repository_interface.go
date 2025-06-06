// repository/target_repository_interface.go
package target

import "uptime-monitor/model"

type TargetRepository interface {
	Add(name, url string, interval int) (model.Target, error)
	List() ([]model.Target, error)
}
