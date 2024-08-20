package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	database "github.com/harry-fruit/simple-go-rest-api/db"
	"github.com/harry-fruit/simple-go-rest-api/models"
)

type EntityRepository struct {
	db *database.SQLDatabase
}

func NewEntityRepository(db *database.SQLDatabase) *EntityRepository {
	return &EntityRepository{db: db}
}

func (er EntityRepository) FindById(id int) *models.Entity {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := er.db.QueryContext(ctx, "SELECT id, unique_code, description FROM entities WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var entity models.Entity

	if rows.Next() {
		err := rows.Scan(&entity.ID, &entity.UniqueCode, &entity.Description)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		return nil
	}

	return &entity
}

func (er EntityRepository) FindByUniqueCode(uniqueCode string) *models.Entity {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := er.db.QueryContext(ctx, "SELECT id, unique_code, description FROM entities WHERE unique_code = ?", uniqueCode)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var entity models.Entity

	if rows.Next() {
		err := rows.Scan(&entity.ID, &entity.UniqueCode, &entity.Description)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		return nil
	}

	return &entity
}

func (er EntityRepository) Create(entity *models.Entity) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := er.db.ExecContext(ctx, "INSERT INTO entities (unique_code, description) VALUES (?, ?)", entity.UniqueCode, entity.Description)

	if err != nil {
		return fmt.Errorf("error creating entity: %v", err)
	}

	return nil
}

func (er EntityRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := er.db.ExecContext(ctx, "DELETE FROM entities WHERE id = ?", id)

	if err != nil {
		return fmt.Errorf("error deleting entity: %v", err)
	}

	return nil
}

func (er EntityRepository) Update(payload map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id int = payload["id"].(int)
	delete(payload, "id")

	updateStatement := "UPDATE entities "
	setStatement := "SET "
	whereStatement := " WHERE id = ?"
	values := []interface{}{}

	index := 0
	for key, value := range payload {
		setStatement += key + " = ?"

		if index < len(payload)-1 {
			setStatement += ", "
		}

		values = append(values, value)
		index++
	}

	values = append(values, id)

	_, err := er.db.ExecContext(ctx, updateStatement+setStatement+whereStatement, values...)

	if err != nil {
		return fmt.Errorf("error while updating entity: %v", err)
	}

	return nil
}
