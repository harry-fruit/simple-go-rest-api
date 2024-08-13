package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() *sql.DB {
	var err error
	database, err := sql.Open("sqlite3", "./mydb.db")

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer database.Close()

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS people (firstname TEXT, lastname TEXT)")

	if err != nil {
		log.Fatalf("Error preparing SQL statement: %v", err)
	}

	defer statement.Close()

	_, err = statement.Exec()

	if err != nil {
		log.Fatalf("Error executing SQL statement: %v", err)
	}

	return database
}
