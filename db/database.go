package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLDatabase struct {
	*sql.DB
}

func newSQLiteDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "../db/db.db")

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	return db
}

func NewSQLDatabase(SGBD string) *SQLDatabase {
	var db *sql.DB
	var sqlDatabase *SQLDatabase

	switch SGBD {
	case "sqlite3":
		db = newSQLiteDatabase()
	default:
		log.Fatalf("There's no implementation for the SGBD %s", SGBD)
	}

	sqlDatabase = &SQLDatabase{db}

	return sqlDatabase
}
