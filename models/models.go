package models

type Vehicle struct {
	Id    int    `json:"id,omitempty" db:"id"`
	Make  string `json:"make" db:"make"`
	Model string `json:"model" db:"model"`
	Price int    `json:"price" db:"price"`
}
