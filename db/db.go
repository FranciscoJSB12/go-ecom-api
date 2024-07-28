package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN()) // We just opened, but we still need to connect

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
