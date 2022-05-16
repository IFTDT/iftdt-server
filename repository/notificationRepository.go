package repository

import (
	"strconv"

	"github.com/iftdt/server/entity"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(notification entity.Notification) entity.Notification
	FindOne(id string) entity.Notification
	FindAll() []entity.Notification
}

type notificationRepository struct {
	connection *gorm.DB
}

func NewNotificationRepository(connection *gorm.DB) NotificationRepository {
	return &notificationRepository{
		connection: connection,
	}
}

func (db *notificationRepository) Create(notification entity.Notification) entity.Notification {
	db.connection.Create(&notification)
	return db.FindOne(string(strconv.Itoa(int(notification.ID))))
}

func (db *notificationRepository) FindOne(id string) entity.Notification {
	var notification entity.Notification
	db.connection.Preload("User").Where("id=?", id).First(&notification)
	return notification
}

func (db *notificationRepository) FindAll() []entity.Notification {
	var notification []entity.Notification
	db.connection.Preload("User").Find(&notification)
	return notification
}
