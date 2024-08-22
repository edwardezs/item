package repo

import (
	"test/internal/model"

	"github.com/jmoiron/sqlx"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) Create(item model.Item) (int, error) {
	var itemId int
	err := r.db.QueryRow(createItemQuery, item.Title, item.Description, item.Done).Scan(&itemId)
	if err != nil {
		return 0, err
	}

	return itemId, nil
}

func (r *ItemPostgres) GetAll() ([]model.Item, error) {
	var items []model.Item
	if err := r.db.Select(&items, selectItemsQuery); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ItemPostgres) GetById(itemId int) (model.Item, error) {
	var item model.Item
	if err := r.db.Get(&item, selectItemQuery, itemId); err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (r *ItemPostgres) Delete(itemId int) (string, error) {
	var title string
	if err := r.db.QueryRow(deleteItemQuery, itemId).Scan(&title); err != nil {
		return "", err
	}
	return title, nil
}
