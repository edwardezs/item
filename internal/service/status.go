package service

import (
	"test/internal/model"
	"test/internal/repo"
)

type StatusService struct {
	repo repo.Status
}

func NewStatusService(repo repo.Status) *StatusService {
	return &StatusService{repo: repo}
}

func (s *StatusService) GetStatus() ([]model.Status, error) {
	return s.repo.GetAll()
}

func (s *StatusService) ChangeStatus(status bool) (bool, error) {
	return s.repo.ChangeAll(status)
}
