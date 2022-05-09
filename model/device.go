package model

import (
	"github.com/iftdt/server/dao"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	Device string `json:"device" gorm:"unique"`
}

func init() {
	dao.DB.AutoMigrate(&Device{})
}

func Create(device *Device) (err error) {
	err = dao.DB.Create(device).Error
	return
}

func FindOne(id string) (device *Device, err error) {
	if err = dao.DB.Where("id=?", id).First(&device).Error; err != nil {
		return nil, err
	}
	return device, nil
}

func FindAll() (devices []Device, err error) {
	if err = dao.DB.Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}

func Update(device *Device) (err error) {
	err = dao.DB.Save(device).Error
	return
}
