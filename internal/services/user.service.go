package services

import (
	"fmt"

	database "github.com/harry-fruit/simple-go-rest-api/db"
	"github.com/harry-fruit/simple-go-rest-api/internal/models"
	"github.com/harry-fruit/simple-go-rest-api/internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(db *database.SQLDatabase) *UserService {
	return &UserService{
		userRepository: repositories.NewUserRepository(db),
	}
}

func (us *UserService) Create(user *models.User) (*models.User, error) {
	existentUser := us.userRepository.FindByLogin(user.Login)

	if existentUser != nil {
		return nil, fmt.Errorf("login '%s' is in use", user.Login)
	}

	err := us.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Name:  user.Name,
		Login: user.Login,
	}

	return newUser, nil
}

func (us *UserService) Delete(id int) error {
	user := us.userRepository.FindById(id)

	if user == nil {
		return fmt.Errorf("user not found")
	}

	error := us.userRepository.Delete(id)

	if error != nil {
		return error
	}

	return nil
}

func (us *UserService) FindById(id int) *models.User {
	return us.userRepository.FindById(id)
}

func (us *UserService) SetPassword(id int, password string) error {
	user := us.userRepository.FindById(id)

	if user == nil {
		return fmt.Errorf("user not found")
	}

	error := us.userRepository.SetPassword(id, password)

	if error != nil {
		return error
	}

	return nil
}

func (us *UserService) Update(userPayload map[string]interface{}) (*models.User, error) {
	userId := userPayload["id"].(int)
	existentUser := us.userRepository.FindById(userId)

	if existentUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	err := us.userRepository.Update(userPayload)

	if err != nil {
		return nil, err
	}

	newUser := us.userRepository.FindById(userId)

	return newUser, nil
}
