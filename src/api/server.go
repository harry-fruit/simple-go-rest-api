package api

import (
	"fmt"
	"net/http"

	"github.com/harry-fruit/simple-go-rest-api/api/controllers"
	appTypes "github.com/harry-fruit/simple-go-rest-api/types"
)

type Server struct {
	addr        string
	controllers []appTypes.Controller
	http.ServeMux
}

func NewServer(addr string) *Server {
	var userController = controllers.NewUserController("/users")

	return &Server{
		addr:     addr,
		ServeMux: *http.NewServeMux(),
		controllers: []appTypes.Controller{
			userController.Controller,
		},
	}
}

func (s *Server) Start() error {
	fmt.Println("Server is starting on ", s.addr)
	return http.ListenAndServe(":8080", &s.ServeMux)
}

func (s *Server) AddControllers() {
	for _, controller := range s.controllers {
		controller.SetRoutes(&s.ServeMux)
	}
}
