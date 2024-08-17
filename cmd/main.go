package main

import (
	"log"

	"github.com/harry-fruit/simple-go-rest-api/api"
	database "github.com/harry-fruit/simple-go-rest-api/db"
)

func main() {
	db := database.NewSQLDatabase("sqlite3")
	server := api.NewServer(":8080", db)

	server.SetControllers()

	log.Fatal(server.Start())
}
