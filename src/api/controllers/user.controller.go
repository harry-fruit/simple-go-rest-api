package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/harry-fruit/simple-go-rest-api/api/services"
	database "github.com/harry-fruit/simple-go-rest-api/db"
	"github.com/harry-fruit/simple-go-rest-api/types"
)

type UserController struct {
	types.Controller
	db *database.SQLDatabase
}

func NewUserController(basePath string, db *database.SQLDatabase) *UserController {
	userController := &UserController{
		db: db,
		Controller: types.Controller{
			BasePath: basePath,
		},
	}

	userController.setRoutes()

	return userController
}

func (uc *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	userService := services.NewUserService(uc.db)

	user := userService.FindById(1)

	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) setRoutes() {
	uc.Routes = []types.Route{
		{Method: "GET", Path: "/{id}", Handler: uc.FindById},
	}
}
