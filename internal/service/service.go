package service

import (
	"items/internal/model"
	"items/internal/repo"
)

type Item interface {
	Create(item model.Item) (int, error)
	GetAll() ([]model.Item, error)
	GetById(itemId int) (model.Item, error)
	Delete(itemId int) error
}

type Status interface {
	GetStatus() ([]model.Status, error)
}

type Service struct {
	Item
	Status
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Item:   NewItemService(repos.Item),
		Status: NewStatusService(repos.Status),
	}
}
