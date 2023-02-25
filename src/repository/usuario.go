package repository

import (
	"api-go/src/model"
	"database/sql"
)

type Usuario struct {
	db *sql.DB
}

func CreateRepository(db *sql.DB) *Usuario {
	return &Usuario{db}
}

func (u Usuario) Save(usuario model.Usuario) (int, error) {
	statement, err := u.db.Prepare("INSERT INTO usuario (nome, nick, email, senha) VALUES (?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}
