package service

import (
	"github.com/iftdt/server/entity"
	"github.com/iftdt/server/repository"
)

type UserService interface {
	Create(user entity.User) entity.User
	FindOne(id string) entity.User
	FindAll() []entity.User

	Login(user entity.User) (*entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (service *userService) Create(user entity.User) entity.User {
	return service.userRepository.Create(user)
}

func (service *userService) FindOne(id string) entity.User {
	return service.userRepository.FindOne(id)
}

func (service *userService) FindAll() []entity.User {
	return service.userRepository.FindAll()
}

func (service *userService) Login(user entity.User) (*entity.User, error) {
	return service.userRepository.Login(user)
}
