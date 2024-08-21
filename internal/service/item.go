package service

import (
	"todo/internal/model"
	"todo/internal/repo"
)

type ItemService struct {
	repo repo.Item
}

func NewItemService(repo repo.Item) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) Create(item model.Item) (int, error) {
	return s.repo.Create(item)
}

func (s *ItemService) GetAll() ([]model.Item, error) {
	return s.repo.GetAll()
}

func (s *ItemService) GetById(itemId int) (model.Item, error) {
	return s.repo.GetById(itemId)
}

func (s *ItemService) Delete(itemId int) error {
	return s.repo.Delete(itemId)
}
