package models

type Entity struct {
	ID          int    `json:"id" db:"id"`
	UniqueCode  string `json:"unique_code" db:"unique_code"`
	Description string `json:"description" db:"description"`
}
