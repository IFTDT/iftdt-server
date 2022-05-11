package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	err := model.Notification{}.Create(&notification)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": notification,
		})
	}
}
