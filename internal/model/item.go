package model

type Item struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description,omitempty" db:"description"`
	Done        bool   `json:"done" db:"done"`
}
