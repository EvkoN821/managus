package interfaces

import "github.com/IlyaZayats/managus/internal/entity"

type ManagerRepository interface {
	GetManagers() ([]entity.Manager, error)
	UpdateManager(manager entity.Manager) error
	InsertManager(manager entity.Manager) error
	DeleteManager(id int) error
}
