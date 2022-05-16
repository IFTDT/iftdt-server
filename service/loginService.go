package service

import (
	"github.com/iftdt/server/entity"
	"github.com/iftdt/server/repository"
)

type LoginService interface {
	Login(loginUser entity.User)
}

type loginService struct {
	userRepository repository.UserRepository
}

func NewLoginService(repo repository.UserRepository) LoginService {
	return &loginService{
		userRepository: repo,
	}
}

func (service *loginService) Login(loginUser entity.User) {

}
