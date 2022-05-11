package model

import (
	"time"

	"github.com/iftdt/server/dao"
)

type Device struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Device    string    `json:"device" gorm:"unique"`

	UserId uint `json:"user_id"`
	User   User `json:"user"`
}

func init() {
	dao.DB.AutoMigrate(&Device{})
}

func (d Device) Create(device *Device) (err error) {
	err = dao.DB.Create(device).Error
	return
}

func (d Device) FindOne(id string) (device *Device, err error) {
	if err = dao.DB.Preload("User").Where("id=?", id).First(&device).Error; err != nil {
		return nil, err
	}
	return device, nil
}

func (d Device) FindAll() (devices []Device, err error) {
	if err = dao.DB.Preload("User").Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}

func (d Device) Update(device *Device) (err error) {
	err = dao.DB.Save(device).Error
	return
}
