package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLDatabase struct {
	*sql.DB
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
	sqlDatabase.migrate()

	return sqlDatabase
}

func newSQLiteDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/mydb.db")

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	return db
}

func (db *SQLDatabase) migrate() {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL,
		login TEXT UNIQUE NOT NULL,
		password TEXT
	);
	`

	_, err := db.Exec(sqlStmt)

	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
	fmt.Println("Migration has been executed")
}
