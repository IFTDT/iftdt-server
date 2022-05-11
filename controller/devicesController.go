package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iftdt/server/model"
)

type DeviceController struct{}

func (device DeviceController) FindAll(c *gin.Context) {
	devices, err := model.Device{}.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": devices,
		})
	}
}

func (device DeviceController) Create(c *gin.Context) {
	var d model.Device
	c.BindJSON(&d)
	err := model.Device{}.Create(&d)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": d,
		})
	}
}

func (device DeviceController) Update(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	d, err := model.Device{}.FindOne(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&d)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": d,
		})
	}
}
