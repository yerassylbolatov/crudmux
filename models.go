package main

type Vehicle struct {
	Id    int    `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Price int    `json:"price"`
}
