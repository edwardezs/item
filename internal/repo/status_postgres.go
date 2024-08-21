package repo

import (
	"todo/internal/model"

	"github.com/jmoiron/sqlx"
)

type StatusPostgres struct {
	db *sqlx.DB
}

func NewStatusPostgres(db *sqlx.DB) *StatusPostgres {
	return &StatusPostgres{db: db}
}

func (r *StatusPostgres) GetAll() ([]model.Status, error) {
	var tableStatuses []model.Status
	if err := r.db.Select(&tableStatuses, selectStatusQuery); err != nil {
		return nil, err
	}

	return tableStatuses, nil
}
