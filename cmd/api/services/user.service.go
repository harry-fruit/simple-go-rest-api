package services

import (
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
