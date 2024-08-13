package main

// import (
// 	"log"

// 	"github.com/harry-fruit/simple-go-rest-api/api"
// )

import (
	database "github.com/harry-fruit/simple-go-rest-api/db"
)

func main() {
	database.InitDatabase()
	// user := entities.UserEntity.Get(1)
	// server := api.NewServer(":8080")
	// log.Fatal(server.Start())
}
