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

type EntityController struct {
	types.Controller
	db *database.SQLDatabase
}

func NewEntityController(basePath string, db *database.SQLDatabase) *EntityController {
	entityController := &EntityController{
		db: db,
		Controller: types.Controller{
			BasePath: basePath,
		},
	}

	entityController.setRoutes()

	return entityController
}

func (ec *EntityController) Create(w http.ResponseWriter, r *http.Request) {
	var newEntity models.Entity
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newEntity)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if newEntity.UniqueCode == "" {
		http.Error(w, "'UniqueCode' is required", http.StatusBadRequest)
		return
	}

	entityService := services.NewEntityService(ec.db)
	entity, err := entityService.Create(&newEntity)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) //TODO: Tratar tipos de erros diferentes
		return
	}

	json.NewEncoder(w).Encode(entity)

}

func (ec *EntityController) Delete(w http.ResponseWriter, r *http.Request) {
	entityService := services.NewEntityService(ec.db)

	idInString := r.PathValue("id")
	id, err := strconv.Atoi(idInString)

	//TODO: Refact -- Validate input
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	err = entityService.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ec *EntityController) Update(w http.ResponseWriter, r *http.Request) {
	entityService := services.NewEntityService(ec.db)

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

	_, uniqueCodeOk := payload["unique_code"].(string)
	_, descriptionOk := payload["description"].(string)

	if payloadUniqueCode, ok := payload["unique_code"]; ok && payloadUniqueCode == nil {
		http.Error(w, "unique_code can't be null", http.StatusBadRequest)
		return
	}

	if !uniqueCodeOk && !descriptionOk {
		http.Error(w, "Invalid or missing 'unique_code' and 'description'", http.StatusBadRequest)
		return
	}

	payload["id"] = id
	entity, err := entityService.Update(payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(entity)
}

func (ec *EntityController) FindById(w http.ResponseWriter, r *http.Request) {
	entityService := services.NewEntityService(ec.db)

	idInString := r.PathValue("id")
	id, err := strconv.Atoi(idInString)

	//TODO: Refact -- Validate input
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID"))
		return
	}

	entity := entityService.FindById(id)
	json.NewEncoder(w).Encode(entity)
}

func (ec *EntityController) setRoutes() {
	ec.Routes = []types.Route{
		{Method: "GET", Path: "/{id}", Handler: ec.FindById},
		{Method: "POST", Path: "/", Handler: ec.Create},
		{Method: "DELETE", Path: "/{id}", Handler: ec.Delete},
		{Method: "PATCH", Path: "/{id}", Handler: ec.Update},
	}
}