package main

import (
	"log"

	"github.com/harry-fruit/simple-go-rest-api/api"
)

func main() {

	// mux := http.NewServeMux()
	// // sqlDatabase := database.NewSQLDatabase("sqlite3")

	// mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	json.NewEncoder(w).Encode(&HelloWorld{Hello: "world"})
	// })

	// mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
	// 	json.NewEncoder(w).Encode(&HelloWorld{Hello: "world - Post"})
	// })

	// http.ListenAndServe(":8080", mux)
	// // userRepo := repositories.NewUserRepository(sqlDatabase)
	// // fmt.Println(userRepo.FindById(1))
	// // user := entities.UserEntity.Get(1)
	// // server := api.NewServer(":8080")
	// // log.Fatal(server.Start())

	server := api.NewServer(":8080")
	// sqlDatabase := database.NewSQLDatabase("sqlite3")

	// userRepo := repositories.NewUserRepository(sqlDatabase)
	// fmt.Println(userRepo.FindById(1))
	// user := entities.UserEntity.Get(1)
	// server := api.NewServer(":8080")
	server.AddControllers()

	log.Fatal(server.Start())
}
