package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	dto "github.com/harry-fruit/simple-go-rest-api/internal/dtos"
	"github.com/harry-fruit/simple-go-rest-api/internal/types"
	httpUtil "github.com/harry-fruit/simple-go-rest-api/internal/utils/http"
)

func (uc *UserController) getHandlers() []types.Route {
	return []types.Route{
		{Method: "GET", Path: "/{id}", Handler: uc.FindById},
		{Method: "POST", Path: "/", Handler: uc.Create},
		{Method: "POST", Path: "/{id}/set-password", Handler: uc.SetPassword},
		{Method: "DELETE", Path: "/{id}", Handler: uc.Delete},
		{Method: "PATCH", Path: "/{id}", Handler: uc.Update},
	}
}

// @Summary Find user
// @Description Find user by ID
// @Tags Users
// @Produce  json
// @Param  id  path  int  true  "user's id"
// @Success 200 {object} httpUtil.HTTPResponse{data=models.User} "user found"
// @Failure 400 {object} httpUtil.HTTPResponse{data=nil} "bad request"
// @Failure 404 {object} httpUtil.HTTPResponse{data=nil} "user not found"
// @Router /users/{id} [get]
func (uc *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idInString := vars["id"]
	id, err := strconv.Atoi(idInString)
	fmt.Println(idInString)

	//TODO: Refact -- Validate input
	if err != nil {
		httpError := &httpUtil.HTTPError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  httpUtil.BadPayload,
		}
		httpError.SendError(w)
		return
	}

	user := uc.userService.FindById(id)

	if user == nil {
		httpError := &httpUtil.HTTPError{
			StatusCode: http.StatusBadRequest,
			Message:    "user not found",
			ErrorType:  httpUtil.BadPayload,
		}
		httpError.SendError(w)
		return
	}

	httpResponse := &httpUtil.HTTPResponse{
		StatusCode: http.StatusOK,
		Message:    "user found",
		Data:       user,
	}

	httpResponse.Send(w)
}

// @Summary Create a new user
// @Description Create a new user then return it
// @Tags Users
// @Accept  json
// @Produce  json
// @Param  user  body  dtos.UserPayloadDTO  true  "user payload"
// @Success 201 {object} httpUtil.HTTPResponse{data=models.User} "user created"
// @Failure 400 {object} httpUtil.HTTPResponse{data=nil} "bad request"
// @Router /users/ [post]
func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var userPayload dto.UserPayloadDTO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userPayload)

	//TODO: Refact -- Error Handler
	if err != nil {
		httpError := &httpUtil.HTTPError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  httpUtil.BadPayload,
		}
		httpError.SendError(w)
		return
	}

	user, httpError := uc.userService.Create(&userPayload)

	if httpError != nil {
		httpError.SendError(w)
		return
	}

	httpResponse := &httpUtil.HTTPResponse{
		StatusCode: http.StatusCreated,
		Message:    "user created",
		Data:       user,
	}

	httpResponse.Send(w)
}

// @Summary Set user's password
// @Description Set user's password by ID
// @Tags Users
// @Produce  json
// @Param  id  path  int  true  "user's id"
// @Param  password  body  dtos.SetUserPasswordPayloadDTO  true  "user's password"
// @Success 204 "user's password set"
// @Failure 400 {object} httpUtil.HTTPResponse{data=nil} "bad request"
// @Failure 404 {object} httpUtil.HTTPResponse{data=nil} "user not found"
// @Router /users/{id}/set-password [post]
func (uc *UserController) SetPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idInString := vars["id"]
	id, err := strconv.Atoi(idInString)

	//TODO: Refact -- Validate input
	if err != nil {
		httpResponse := &httpUtil.HTTPResponse{
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
		httpResponse := &httpUtil.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "bad request",
			Data:       nil,
		}

		httpResponse.Send(w)
		return
	}

	password, ok := payload["password"].(string)

	if !ok {
		httpResponse := &httpUtil.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid or missing password",
			Data:       nil,
		}

		httpResponse.Send(w)
		return
	}

	err = uc.userService.SetPassword(id, password)

	if err != nil {
		// TODO: Refact -- Handle error
		httpResponse := &httpUtil.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "internal server error",
			Data:       nil,
		}

		httpResponse.Send(w)
		return
	}

	httpResponse := &httpUtil.HTTPResponse{
		StatusCode: http.StatusNoContent,
	}

	httpResponse.Send(w)
}

// @Summary Update user
// @Description Update user then return it
// @Tags Users
// @Accept  json
// @Produce  json
// @Param  id  path  int  true  "user's id"
// @Param  user  body  dtos.UserPayloadDTO  true  "user payload"
// @Success 201 {object} httpUtil.HTTPResponse{data=models.User} "user updated"
// @Failure 400 {object} httpUtil.HTTPResponse{data=nil} "bad request"
// @Failure 404 {object} httpUtil.HTTPResponse{data=nil} "user not found"
// @Router /users/{id} [patch]
func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idInString := vars["id"]
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

	if payloadLogin, ok := payload["login"]; ok && payloadLogin == nil {
		http.Error(w, "login can't be null", http.StatusBadRequest)
		return
	}

	if len(payload) == 0 {
		http.Error(w, "payload can't be empty", http.StatusBadRequest)
		return
	}

	payload["id"] = id
	user, err := uc.userService.Update(payload)

	if err != nil {
		// TODO: Refact -- Handle error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// @Summary Delete user
// @Description Delete user by ID
// @Tags Users
// @Produce  json
// @Param  id  path  int  true  "user's id"
// @Success 204 "user deleted"
// @Failure 400 {object} httpUtil.HTTPResponse{data=nil} "bad request"
// @Failure 404 {object} httpUtil.HTTPResponse{data=nil} "user not found"
// @Router /users/{id} [delete]
func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idInString := vars["id"]
	id, err := strconv.Atoi(idInString)

	//TODO: Refact -- Validate input
	if err != nil {
		httpResponse := &httpUtil.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "invalid id",
			Data:       nil,
		}

		httpResponse.Send(w)
		return
	}

	err = uc.userService.Delete(id)

	if err != nil {
		// TODO: Refact -- Handle error
		httpResponse := &httpUtil.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "bad request",
			Data:       nil,
		}

		httpResponse.Send(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
