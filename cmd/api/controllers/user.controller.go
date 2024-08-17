package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/harry-fruit/simple-go-rest-api/api/services"
	database "github.com/harry-fruit/simple-go-rest-api/db"
	"github.com/harry-fruit/simple-go-rest-api/models"
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

	idInString := r.PathValue("id")
	id, err := strconv.Atoi(idInString)

	//TODO: Refact -- Validate input
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	user := userService.FindById(id)
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	userService := services.NewUserService(uc.db)

	user, error := userService.Create(&newUser)

	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest) //TODO: Tratar tipos de erros diferentes
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) SetPassword(w http.ResponseWriter, r *http.Request) {
	userService := services.NewUserService(uc.db)

	idInString := r.PathValue("id")
	id, err := strconv.Atoi(idInString)

	//TODO: Refact -- Validate input
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	var payload map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&payload)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	password, ok := payload["password"].(string)

	if !ok {
		http.Error(w, "Invalid or missing password", http.StatusBadRequest)
		return
	}

	err = userService.SetPassword(id, password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	userService := services.NewUserService(uc.db)

	idInString := r.PathValue("id")
	id, err := strconv.Atoi(idInString)

	//TODO: Refact -- Validate input
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	err = userService.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	userService := services.NewUserService(uc.db)

	idInString := r.PathValue("id")
	id, err := strconv.Atoi(idInString)

	//TODO: Refact -- Validate input
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	var payload map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&payload)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	_, nameOk := payload["name"].(string)
	_, loginOk := payload["login"].(string)

	if payloadLogin, ok := payload["login"]; ok && payloadLogin == nil {
		http.Error(w, "login can't be null", http.StatusBadRequest)
		return
	}

	if !nameOk && !loginOk {
		http.Error(w, "Invalid or missing name and login", http.StatusBadRequest)
		return
	}

	payload["id"] = id
	user, err := userService.Update(payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) setRoutes() {
	uc.Routes = []types.Route{
		{Method: "GET", Path: "/{id}", Handler: uc.FindById},
		{Method: "POST", Path: "/", Handler: uc.Create},
		{Method: "POST", Path: "/{id}/set-password", Handler: uc.SetPassword},
		{Method: "DELETE", Path: "/{id}", Handler: uc.Delete},
		{Method: "PATCH", Path: "/{id}", Handler: uc.Update},
	}
}
