package services

import (
	"github.com/IlyaZayats/managus/internal/entity"
	"github.com/IlyaZayats/managus/internal/interfaces"
	"strconv"
)

type ManagerService struct {
	repo interfaces.ManagerRepository
}

func NewManagerService(repo interfaces.ManagerRepository) (*ManagerService, error) {
	return &ManagerService{
		repo: repo,
	}, nil
}

func (s *ManagerService) GetManagers() ([]map[string]string, error) {
	managers, err := s.repo.GetManagers()
	if err != nil {
		return nil, err
	}
	managersSlice := []map[string]string{}
	for _, item := range managers {
		facultiesMap := map[string]string{
			"id":   strconv.Itoa(item.Id),
			"name": item.Name,
		}
		managersSlice = append(managersSlice, facultiesMap)
	}
	return managersSlice, nil
}

func (s *ManagerService) InsertManager(name string) error {
	return s.repo.InsertManager(entity.Manager{Id: 0, Name: name})
}

func (s *ManagerService) UpdateManagers(id int, name string) error {
	return s.repo.UpdateManager(entity.Manager{Id: id, Name: name})
}

func (s *ManagerService) DeleteManager(id int) error {
	return s.repo.DeleteManager(id)
}
