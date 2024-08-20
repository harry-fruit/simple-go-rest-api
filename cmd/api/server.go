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

func getControllers(db *database.SQLDatabase) []appTypes.Controller {
	var userController = controllers.NewUserController("/users", db)
	var entityController = controllers.NewEntityController("/entities", db)

	return []appTypes.Controller{
		userController.Controller,
		entityController.Controller,
	}
}

func NewServer(addr string, db *database.SQLDatabase) *Server {
	return &Server{
		addr:        addr,
		ServeMux:    *http.NewServeMux(),
		db:          db,
		controllers: getControllers(db),
	}
}

func (s *Server) SetControllers() {
	fmt.Println("----- Setting controllers... -----")
	for _, controller := range s.controllers {
		controller.Init(&s.ServeMux)
	}
	fmt.Println("----- Controllers set -----")
}

func (s *Server) Start() error {
	fmt.Println("Server is starting on ", s.addr)
	return http.ListenAndServe(":8080", &s.ServeMux)
}
