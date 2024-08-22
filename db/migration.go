package db

import "fmt"

type MigrationTable struct {
	Name    string
	SQLStmt string
}

func getTables() []MigrationTable {
	var tables = []MigrationTable{
		{
			Name: "users",
			SQLStmt: `CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			name TEXT NOT NULL,
			login TEXT UNIQUE NOT NULL,
			password TEXT
			);`,
		},
		{
			Name: "entities",
			SQLStmt: `	CREATE TABLE IF NOT EXISTS entities (
			id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
			unique_code TEXT UNIQUE NOT NULL,
			description TEXT
			);`,
		},
	}

	return tables
}

func (db *SQLDatabase) migrate() {
	fmt.Println("----- Migrating database... -----")
	tables := getTables()

	for _, table := range tables {
		_, err := db.Exec(table.SQLStmt)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Table '%s' created\n", table.Name)
	}
	fmt.Println("----- Database migrated -----")
}
