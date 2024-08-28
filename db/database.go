package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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
	case "postgres":
		db = newPostgresDatabase()
	default:
		log.Fatalf("There's no implementation for the SGBD %s", SGBD)
	}

	sqlDatabase = &SQLDatabase{db}

	db.Exec("SELECT 1=1")

	return sqlDatabase
}

func newPostgresDatabase() *sql.DB {
	connStr := "user=admin password=admin dbname=SGRA sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	// defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Connected to PostgreSQL database!")
	return db
}
