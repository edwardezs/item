package service

import (
	"items/internal/model"
	"items/internal/repo"
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
