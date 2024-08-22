package repo

import (
	"test/internal/model"

	"github.com/jmoiron/sqlx"
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
	GetAll() ([]model.Status, error)
	ChangeAll(status bool) (bool, error)
}

type Repository struct {
	Item
	User
	Status
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Item:   NewItemPostgres(db),
		User:   NewUserPostgres(db),
		Status: NewStatusPostgres(db),
	}
}
