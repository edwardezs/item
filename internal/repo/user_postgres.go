package repo

import (
	"items/internal/model"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user model.User) (int, error) {
	var itemId int
	err := r.db.QueryRow(createUserQuery, user.Name, user.Job).Scan(&itemId)
	if err != nil {
		return 0, err
	}

	return itemId, nil
}

func (r *UserPostgres) GetAll() ([]model.User, error) {
	var users []model.User
	if err := r.db.Select(&users, selectUsersQuery); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserPostgres) GetById(userId int) (model.User, error) {
	var user model.User
	if err := r.db.Get(&user, selectUserQuery, userId); err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserPostgres) Delete(userId int) error {
	_, err := r.db.Exec(deleteUserQuery, userId)
	return err
}
