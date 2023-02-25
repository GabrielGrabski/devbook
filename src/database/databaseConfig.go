package database

import (
	"api-go/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.Dns)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
	}

	return db, nil
}
