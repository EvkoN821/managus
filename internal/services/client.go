package services

import (
	"github.com/IlyaZayats/managus/internal/entity"
	"github.com/IlyaZayats/managus/internal/interfaces"
	"strconv"
)

type ClientService struct {
	repo interfaces.ClientRepository
}

func NewClientService(repo interfaces.ClientRepository) (*ClientService, error) {
	return &ClientService{
		repo: repo,
	}, nil
}

func (s *ClientService) GetClients() ([]map[string]string, error) {
	courses, err := s.repo.GetClients()
	if err != nil {
		return nil, err
	}
	coursesSlice := []map[string]string{}
	for _, item := range courses {
		coursesMap := map[string]string{
			"id":         strconv.Itoa(item.Id),
			"manager_id": strconv.Itoa(item.ManagerId),
			"name":       item.Name,
		}
		coursesSlice = append(coursesSlice, coursesMap)
	}
	return coursesSlice, nil
}

func (s *ClientService) InsertClient(managerId int, name string) error {
	return s.repo.InsertClient(entity.Client{Id: 0, ManagerId: managerId, Name: name})
}

func (s *ClientService) UpdateClient(id int, name string) error {
	return s.repo.UpdateClient(entity.Client{Id: id, ManagerId: 0, Name: name})
}

func (s *ClientService) DeleteClient(id int) error {
	return s.repo.DeleteClient(id)
}
