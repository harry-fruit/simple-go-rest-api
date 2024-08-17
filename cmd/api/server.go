package api

import (
	"fmt"
	"net/http"

	"github.com/harry-fruit/simple-go-rest-api/api/controllers"
	database "github.com/harry-fruit/simple-go-rest-api/db"
	appTypes "github.com/harry-fruit/simple-go-rest-api/types"
)

type Server struct {
	addr        string
	controllers []appTypes.Controller
	db          *database.SQLDatabase
	http.ServeMux
}

func NewServer(addr string, db *database.SQLDatabase) *Server {
	var userController = controllers.NewUserController("/users", db)

	return &Server{
		addr:     addr,
		ServeMux: *http.NewServeMux(),
		db:       db,
		controllers: []appTypes.Controller{
			userController.Controller,
		},
	}
}

func (s *Server) Start() error {
	fmt.Println("Server is starting on ", s.addr)
	return http.ListenAndServe(":8080", &s.ServeMux)
}

func (s *Server) SetControllers() {
	for _, controller := range s.controllers {
		controller.Init(&s.ServeMux)
	}
}
