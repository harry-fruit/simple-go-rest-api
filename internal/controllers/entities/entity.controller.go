package controllers

import (
	database "github.com/harry-fruit/simple-go-rest-api/db"
	"github.com/harry-fruit/simple-go-rest-api/internal/services"
	"github.com/harry-fruit/simple-go-rest-api/internal/types"
)

type EntityController struct {
	types.Controller
	db            *database.SQLDatabase
	entityService *services.EntityService
}

func NewEntityController(basePath string, db *database.SQLDatabase) *EntityController {
	entityController := &EntityController{
		db:            db,
		entityService: services.NewEntityService(db),
		Controller: types.Controller{
			BasePath: basePath,
		},
	}

	entityController.setRoutes()

	return entityController
}

func (ec *EntityController) setRoutes() {
	ec.Routes = ec.getHandlers()
}
