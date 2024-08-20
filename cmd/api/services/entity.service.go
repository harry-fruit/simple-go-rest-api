package services

import (
	"fmt"

	"github.com/harry-fruit/simple-go-rest-api/api/repositories"
	database "github.com/harry-fruit/simple-go-rest-api/db"
	"github.com/harry-fruit/simple-go-rest-api/models"
)

type EntityService struct {
	entityRepository *repositories.EntityRepository
}

func NewEntityService(db *database.SQLDatabase) *EntityService {
	return &EntityService{
		entityRepository: repositories.NewEntityRepository(db),
	}
}

func (es *EntityService) Create(entity *models.Entity) (*models.Entity, error) {
	existentEntity := es.entityRepository.FindByUniqueCode(entity.UniqueCode)

	if existentEntity != nil {
		return nil, fmt.Errorf("entity's unique_code '%s' is in use", entity.UniqueCode)
	}

	err := es.entityRepository.Create(entity)

	if err != nil {
		return nil, err
	}

	created := &models.Entity{
		UniqueCode:  entity.UniqueCode,
		Description: entity.Description,
	}

	return created, nil
}

func (es *EntityService) Delete(id int) error {
	entity := es.entityRepository.FindById(id)

	if entity == nil {
		return fmt.Errorf("entity not found")
	}

	error := es.entityRepository.Delete(id)

	if error != nil {
		return error
	}

	return nil
}

func (es *EntityService) FindById(id int) *models.Entity {
	return es.entityRepository.FindById(id)
}

func (es *EntityService) Update(payload map[string]interface{}) (*models.Entity, error) {
	entityID := payload["id"].(int)
	existentEntity := es.entityRepository.FindById(entityID)

	if existentEntity == nil {
		return nil, fmt.Errorf("entity not found")
	}

	err := es.entityRepository.Update(payload)

	if err != nil {
		return nil, err
	}

	newUser := es.entityRepository.FindById(entityID)

	return newUser, nil
}
