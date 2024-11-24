package db

import (
	"database/sql"

	// _ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

func ConnectDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "donation.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
