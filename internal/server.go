package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/harry-fruit/simple-go-rest-api/api"
	database "github.com/harry-fruit/simple-go-rest-api/db"
	entitiesController "github.com/harry-fruit/simple-go-rest-api/internal/controllers/entities"
	usersController "github.com/harry-fruit/simple-go-rest-api/internal/controllers/users"
	appTypes "github.com/harry-fruit/simple-go-rest-api/internal/types"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	addr        string
	controllers []appTypes.Controller
	db          *database.SQLDatabase
	*mux.Router
}

func getControllers(db *database.SQLDatabase) []appTypes.Controller {
	var userController = usersController.NewUserController("/users", db)
	var entityController = entitiesController.NewEntityController("/entities", db)

	return []appTypes.Controller{
		userController.Controller,
		entityController.Controller,
	}
}

func NewServer(addr string, db *database.SQLDatabase) *Server {
	mux := mux.NewRouter()

	return &Server{
		addr:        addr,
		Router:      mux,
		db:          db,
		controllers: getControllers(db),
	}
}

func (s *Server) setControllers() {
	fmt.Println("----- Setting controllers... -----")

	for _, controller := range s.controllers {
		controller.SetRoutes(s.Router)
	}

	s.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	fmt.Println("----- Controllers set -----")
}

func (s *Server) Start() error {
	fmt.Println("Server is starting on ", s.addr)
	s.setControllers()
	return http.ListenAndServe(":8080", s.Router)
}
