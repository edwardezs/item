package repo

import (
	"test/internal/model"

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

func (r *StatusPostgres) ChangeAll(status bool) (bool, error) {
	var newStatus bool
	err := r.db.QueryRow(changeStatusQuery, status).Scan(&newStatus)
	if err != nil {
		return false, err
	}

	return newStatus, nil
}
