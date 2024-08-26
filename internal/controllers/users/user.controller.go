package controllers

import (
	database "github.com/harry-fruit/simple-go-rest-api/db"
	"github.com/harry-fruit/simple-go-rest-api/internal/services"
	"github.com/harry-fruit/simple-go-rest-api/internal/types"
)

type UserController struct {
	types.Controller
	db          *database.SQLDatabase
	userService *services.UserService
}

func NewUserController(basePath string, db *database.SQLDatabase) *UserController {

	userController := &UserController{
		db:          db,
		userService: services.NewUserService(db),
		Controller: types.Controller{
			BasePath: basePath,
		},
	}

	userController.setRoutes()

	return userController
}

func (uc *UserController) setRoutes() {
	uc.Routes = uc.getHandlers()
}
