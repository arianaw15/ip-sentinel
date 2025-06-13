package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewDBConnection(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
		return nil, err
	}

	return db, nil
}
