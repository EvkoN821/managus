package interfaces

import "github.com/IlyaZayats/managus/internal/entity"

type ClientRepository interface {
	GetClients() ([]entity.Client, error)
	UpdateClient(client entity.Client) error
	InsertClient(client entity.Client) error
	DeleteClient(id int) error
}
