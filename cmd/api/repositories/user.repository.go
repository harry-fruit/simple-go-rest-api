package repositories

import (
	"context"
	"log"
	"time"

	database "github.com/harry-fruit/simple-go-rest-api/db"
	"github.com/harry-fruit/simple-go-rest-api/models"
)

type UserRepository struct {
	db *database.SQLDatabase
}

func NewUserRepository(db *database.SQLDatabase) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) FindById(id int) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := u.db.QueryContext(ctx, "SELECT id, name, login FROM users WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var user models.User

	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Login)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		return nil
	}

	return &user
}
