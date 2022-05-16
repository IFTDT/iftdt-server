package service

import (
	"strconv"

	"github.com/iftdt/server/common"
	"github.com/iftdt/server/entity"
	"github.com/iftdt/server/repository"
)

type NotificationService interface {
	Create(notification entity.Notification) entity.Notification
	FindOne(id string) entity.Notification
	FindAll() []entity.Notification
}

type notificationService struct {
	repo       repository.NotificationRepository
	deviceRepo repository.DeviceRepository
}

func NewNotificationService(repo repository.NotificationRepository, deviceRepo repository.DeviceRepository) NotificationService {
	return &notificationService{
		repo:       repo,
		deviceRepo: deviceRepo,
	}
}

func (service *notificationService) Create(notification entity.Notification) entity.Notification {
	var user_id = string(strconv.Itoa(int(notification.UserID)))
	device, err := service.deviceRepo.FindOneByUserId(user_id)
	if err != nil {
		panic("not found device when push notification")
	}
	go common.PushNotifcation(device.DeviceToken, notification.Payload)
	return service.repo.Create(notification)
}

func (service *notificationService) FindOne(id string) entity.Notification {
	return service.repo.FindOne(id)
}

func (service *notificationService) FindAll() []entity.Notification {
	return service.repo.FindAll()
}
