package repositories

import (
	"context"
	"fmt"
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

func (u UserRepository) FindByLogin(login string) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := u.db.QueryContext(ctx, "SELECT id, name, login FROM users WHERE login = ?", login)

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

func (u UserRepository) Create(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := u.db.ExecContext(ctx, "INSERT INTO users (name, login) VALUES (?, ?)", user.Name, user.Login)

	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

func (u UserRepository) SetPassword(id int, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := u.db.ExecContext(ctx, "UPDATE users SET password = ? WHERE id = ?", password, id)

	if err != nil {
		return fmt.Errorf("error setting password: %v", err)
	}

	return nil
}

// func (u UserRepository) Update() error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	_, err := u.db.ExecContext(ctx, "UPDATE users SET password = ? WHERE id = ?", password, id)

// 	if err != nil {
// 		return fmt.Errorf("error setting password: %v", err)
// 	}

// 	return nil
// }
