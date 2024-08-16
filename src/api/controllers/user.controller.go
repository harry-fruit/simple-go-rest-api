package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/harry-fruit/simple-go-rest-api/models"
	"github.com/harry-fruit/simple-go-rest-api/types"
)

type UserController struct {
	types.Controller
}

var routes = []types.Route{
	{Method: "GET", Path: "/{id}", Handler: FindById},
}

func NewUserController(basePath string) *UserController {
	return &UserController{
		Controller: types.Controller{
			BasePath: basePath,
			Routes:   routes,
		},
	}
}

func FindById(w http.ResponseWriter, r *http.Request) {
	user := &models.User{
		ID:    1,
		Login: "harryfruit",
		Name:  "Harry Fruit",
	}

	json.NewEncoder(w).Encode(user)
}
