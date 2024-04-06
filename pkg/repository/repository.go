package repository

import (
	"microservice"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user microservice.User) (int, error)
	GetUser(username, passwordHash string) (microservice.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
