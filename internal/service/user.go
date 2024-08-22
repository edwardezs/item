package service

import (
	"test/internal/model"
	"test/internal/repo"
)

type UserService struct {
	repo repo.User
}

func NewUserService(repo repo.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user model.User) (int, error) {
	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(userId int) (model.User, error) {
	return s.repo.GetById(userId)
}

func (s *UserService) Delete(userId int) (string, error) {
	return s.repo.Delete(userId)
}
