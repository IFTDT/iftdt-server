package model

import (
	"time"

	"github.com/iftdt/server/dao"
)

type Notification struct {
	ID        uint                   `json:"id"`
	PayLoad   map[string]interface{} `josn:"payload"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`

	UserID uint `json:"user_id"`
	User   User `json:"user"`
}

func init() {
	dao.DB.AutoMigrate(&Notification{})
}

func (n Notification) Create(notification *Notification) (err error) {
	err = dao.DB.Create(notification).Error
	return
}

func (u Notification) FindOne(id string) (notification *Notification, err error) {
	if err = dao.DB.Preload("User").Where("id=?", id).First(&notification).Error; err != nil {
		return nil, err
	}
	return notification, nil
}

func (u Notification) FindAll() (notification []Notification, err error) {
	if err = dao.DB.Preload("User").Find(&notification).Error; err != nil {
		return nil, err
	}
	return notification, nil
}
