package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	database "github.com/harry-fruit/simple-go-rest-api/db"
	"github.com/harry-fruit/simple-go-rest-api/internal/models"
	"github.com/harry-fruit/simple-go-rest-api/internal/services"
	"github.com/harry-fruit/simple-go-rest-api/internal/types"
	response "github.com/harry-fruit/simple-go-rest-api/internal/types/http"
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
	uc.Routes = []types.Route{
		{Method: "GET", Path: "/{id}", Handler: uc.FindById},
		{Method: "POST", Path: "/", Handler: uc.Create},
		{Method: "POST", Path: "/{id}/set-password", Handler: uc.SetPassword},
		{Method: "DELETE", Path: "/{id}", Handler: uc.Delete},
		{Method: "PATCH", Path: "/{id}", Handler: uc.Update},
	}
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	user, error := uc.userService.Create(&newUser)

	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest) //TODO: Tratar tipos de erros diferentes
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	idInString := r.PathValue("id")
	id, err := strconv.Atoi(idInString)

	//TODO: Refact -- Validate input
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	err = uc.userService.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) FindById(w http.ResponseWriter, r *http.Request) {

	idInString := r.PathValue("id")
	id, err := strconv.Atoi(idInString)

	//TODO: Refact -- Validate input
	if err != nil {
		httpResponse := &response.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
			Data:       nil,
		}

		httpResponse.Send(w)
		return
	}

	user := uc.userService.FindById(id)

	if user == nil {
		httpResponse := &response.HTTPResponse{
			StatusCode: http.StatusNotFound,
			Message:    "user not found",
			Data:       nil,
		}

		httpResponse.Send(w)
		return
	}

	httpResponse := &response.HTTPResponse{
		StatusCode: http.StatusOK,
		Message:    "user found",
		Data:       user,
	}

	httpResponse.Send(w)
}

func (uc *UserController) SetPassword(w http.ResponseWriter, r *http.Request) {
	idInString := r.PathValue("id")
	id, err := strconv.Atoi(idInString)

	//TODO: Refact -- Validate input
	if err != nil {
		httpResponse := &response.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
			Data:       nil,
		}

		httpResponse.Send(w)
		return
	}

	var payload map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&payload)

	if err != nil {
		httpResponse := &response.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "bad request",
			Data:       nil,
		}

		httpResponse.Send(w)
		return
	}

	password, ok := payload["password"].(string)

	if !ok {
		httpResponse := &response.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid or missing password",
			Data:       nil,
		}

		httpResponse.Send(w)
		return
	}

	err = uc.userService.SetPassword(id, password)

	if err != nil {
		httpResponse := &response.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "internal server error",
			Data:       nil,
		}

		httpResponse.Send(w)
		return
	}

	httpResponse := &response.HTTPResponse{
		StatusCode: http.StatusNoContent,
	}

	httpResponse.Send(w)
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
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
	user, err := uc.userService.Update(payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}
