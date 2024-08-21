package repo

import (
	"items/internal/model"

	"github.com/jmoiron/sqlx"
)

type Item interface {
	Create(item model.Item) (int, error)
	GetAll() ([]model.Item, error)
	GetById(itemId int) (model.Item, error)
	Delete(itemId int) error
}

type Status interface {
	GetAll() ([]model.Status, error)
}

type Repository struct {
	Item
	Status
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Item:   NewItemPostgres(db),
		Status: NewStatusPostgres(db),
	}
}
