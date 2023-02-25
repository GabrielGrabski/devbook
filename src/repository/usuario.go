package repository

import "database/sql"

type Usuario struct {
	db *sql.DB
}

func CreateRepository(db *sql.DB) *Usuario {
	return &Usuario{db}
}
