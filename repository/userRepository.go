package repository

import (
	"errors"

	"github.com/iftdt/server/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entity.User) entity.User
	FindOne(id string) entity.User
	FindAll() []entity.User
	Login(user entity.User) (*entity.User, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) UserRepository {
	return &userRepository{
		connection: connection,
	}
}

func (db *userRepository) Create(user entity.User) entity.User {
	db.connection.Create(&user)
	return user
}

func (db *userRepository) FindOne(id string) entity.User {
	var user entity.User
	db.connection.Preload("User").Where("id=?", id).First(&user)
	return user
}

func (db *userRepository) FindAll() []entity.User {
	var devices []entity.User
	db.connection.Preload("User").Find(&devices)
	return devices
}

func (db *userRepository) Login(loginUser entity.User) (user *entity.User, err error) {
	if err = db.connection.Where("name=?", loginUser.Name).First(&user).Error; err != nil {
		return nil, err
	}

	if user.Password != loginUser.Password {
		return nil, errors.New("账号或密码错误")
	}

	return user, nil
}
