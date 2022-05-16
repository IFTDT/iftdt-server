package repository

import (
	"github.com/iftdt/server/entity"
	"gorm.io/gorm"
)

type DeviceRepository interface {
	Create(device entity.Device) entity.Device
	Update(device entity.Device)
	Delete(device entity.Device)
	FindOne(id string) entity.Device
	FindOneByUserId(id string) (*entity.Device, error)
	FindAll() []entity.Device
}

type deviceRepository struct {
	connection *gorm.DB
}

func NewDeviceRepository(connection *gorm.DB) DeviceRepository {
	return &deviceRepository{
		connection: connection,
	}
}

func (db *deviceRepository) Create(device entity.Device) entity.Device {
	db.connection.Debug().Create(&device)
	return device
}

func (db *deviceRepository) Update(device entity.Device) {
	db.connection.Save(&device)
}

func (db *deviceRepository) Delete(device entity.Device) {
	db.connection.Delete(&device)
}

func (db *deviceRepository) FindOne(id string) entity.Device {
	var device entity.Device
	db.connection.Preload("User").Where("id=?", id).First(&device)
	return device
}

func (db *deviceRepository) FindOneByUserId(id string) (device *entity.Device, err error) {
	if err = db.connection.Preload("User").Where("user_id=?", id).First(&device).Error; err != nil {
		return nil, err
	}
	return device, nil
}

func (db *deviceRepository) FindAll() []entity.Device {
	var devices []entity.Device
	db.connection.Preload("User").Find(&devices)
	return devices
}
