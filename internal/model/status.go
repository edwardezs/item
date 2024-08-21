package model

type Status struct {
	TableName string `json:"table_name" db:"table_name"`
	ReadOnly  bool   `json:"read_only" db:"read_only"`
}
