package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iftdt/server/common"
	"github.com/iftdt/server/model"
)

type NotificationController struct{}

func (controller NotificationController) FindAll(c *gin.Context) {
	notifications, err := model.Notification{}.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": notifications,
		})
	}
}

func (controller NotificationController) Create(c *gin.Context) {
	var notification model.Notification
	c.BindJSON(&notification)
	var user_id = string(strconv.Itoa(int(notification.UserID)))
	device, err := model.Device{}.FindOneByUserId(user_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	go common.PushNotifcation(device.DeviceToken, notification.Payload)
	err = model.Notification{}.Create(&notification)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"id":      notification.ID,
				"payload": notification.Payload,
				"user_id": notification.UserID,
			},
		})
	}
}
