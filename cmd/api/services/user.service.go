package services

import (
	"fmt"

	"github.com/harry-fruit/simple-go-rest-api/api/repositories"
	database "github.com/harry-fruit/simple-go-rest-api/db"
	"github.com/harry-fruit/simple-go-rest-api/models"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(db *database.SQLDatabase) *UserService {
	return &UserService{
		userRepository: repositories.NewUserRepository(db),
	}
}

func (us *UserService) FindById(id int) *models.User {
	return us.userRepository.FindById(id)
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
