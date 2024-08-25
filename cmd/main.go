package main

import (
	"log"

	database "github.com/harry-fruit/simple-go-rest-api/db"
	api "github.com/harry-fruit/simple-go-rest-api/internal"
)

func main() {
	db := database.NewSQLDatabase("sqlite3")
	server := api.NewServer(":8080", db)

	log.Fatal(server.Start())
}
