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

func (ur UserRepository) FindById(id int) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := ur.db.QueryContext(ctx, "SELECT id, name, login FROM users WHERE id = ?", id)

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

func (ur UserRepository) FindByLogin(login string) *models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := ur.db.QueryContext(ctx, "SELECT id, name, login FROM users WHERE login = ?", login)

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

func (ur UserRepository) Create(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := ur.db.ExecContext(ctx, "INSERT INTO users (name, login) VALUES (?, ?)", user.Name, user.Login)

	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

func (ur UserRepository) SetPassword(id int, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := ur.db.ExecContext(ctx, "UPDATE users SET password = ? WHERE id = ?", password, id)

	if err != nil {
		return fmt.Errorf("error setting password: %v", err)
	}

	return nil
}

func (ur UserRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := ur.db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)

	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	return nil
}

func (u UserRepository) Update(userPayload map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id int = userPayload["id"].(int)
	delete(userPayload, "id")

	updateStatement := "UPDATE users "
	setStatement := "SET "
	whereStatement := " WHERE id = ?"
	values := []interface{}{}

	index := 0
	for key, value := range userPayload {
		setStatement += key + " = ?"

		if index < len(userPayload)-1 {
			setStatement += ", "
		}

		values = append(values, value)
		index++
	}

	values = append(values, id)

	_, err := u.db.ExecContext(ctx, updateStatement+setStatement+whereStatement, values...)

	if err != nil {
		return fmt.Errorf("error while updating user: %v", err)
	}

	return nil
}
