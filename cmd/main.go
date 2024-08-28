package main

import (
	"log"

	"github.com/harry-fruit/simple-go-rest-api/config"
	database "github.com/harry-fruit/simple-go-rest-api/db"
	api "github.com/harry-fruit/simple-go-rest-api/internal"
)

func main() {
	args := config.GetArgs()
	db := database.NewSQLDatabase(args["database"])
	server := api.NewServer(":8080", db)

	log.Fatal(server.Start())
}
