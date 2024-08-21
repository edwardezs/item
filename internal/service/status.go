package service

import (
	"todo/internal/model"
	"todo/internal/repo"
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
