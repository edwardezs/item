package model

type User struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Job  string `json:"job,omitempty" db:"job"`
}
