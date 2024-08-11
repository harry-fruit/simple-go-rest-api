package main

import (
	"log"

	"github.com/harry-fruit/simple-go-rest-api/api"
)

func main() {
	server := api.NewServer(":8080")

	log.Fatal(server.Start())
}
