// repository/target_repository_interface.go
package target

import "uptime-monitor/model"

type TargetRepository interface {
	Add(name, url string, interval int) (model.Target, error)
	Update(id, name, url string, interval int) (model.Target, error)
	List() ([]model.TargetListDto, error)
	GetByID(id string) (model.Target, error)
	Delete(id string) error
}
