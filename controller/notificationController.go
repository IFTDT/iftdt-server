package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iftdt/server/entity"
	"github.com/iftdt/server/service"
)

type NotificationController interface {
	FindAll() []entity.Notification
	Create(ctx *gin.Context) entity.Notification
}

type notificationController struct {
	service service.NotificationService
}

func NewNotificationController(service service.NotificationService) NotificationController {
	return &notificationController{
		service: service,
	}
}

func (c *notificationController) FindAll() []entity.Notification {
	return c.service.FindAll()
}

func (c *notificationController) Create(ctx *gin.Context) entity.Notification {
	var notification entity.Notification
	ctx.BindJSON(&notification)
	return c.service.Create(notification)
}
