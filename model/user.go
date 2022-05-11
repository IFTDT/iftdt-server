package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/iftdt/server/dao"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func init() {
	dao.DB.AutoMigrate(&User{})
}

func (u User) Create(user *User) (err error) {
	err = dao.DB.Create(user).Error
	return
}

func (u User) FindOne(id string) (user *User, err error) {
	if err = dao.DB.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u User) FindAll() (users []User, err error) {
	if err = dao.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u User) Update(user *User) (err error) {
	err = dao.DB.Debug().Save(user).Error
	return
}

func (u User) Login(loginUser User) (user *User, err error) {
	if err = dao.DB.Where("name=?", loginUser.Name).First(&user).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	if user.Password != loginUser.Password {
		return nil, errors.New("账号或密码错误")
	}

	return user, nil
}
