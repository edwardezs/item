package service

import (
	"test/internal/model"
	"test/internal/repo"
)

type Item interface {
	Create(item model.Item) (int, error)
	GetAll() ([]model.Item, error)
	GetById(itemId int) (model.Item, error)
	Delete(itemId int) (string, error)
}

type User interface {
	Create(user model.User) (int, error)
	GetAll() ([]model.User, error)
	GetById(userId int) (model.User, error)
	Delete(userId int) (string, error)
}

type Status interface {
	GetStatus() ([]model.Status, error)
	ChangeStatus(status bool) (bool, error)
}

type Service struct {
	Item
	User
	Status
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Item:   NewItemService(repos.Item),
		User:   NewUserService(repos.User),
		Status: NewStatusService(repos.Status),
	}
}
