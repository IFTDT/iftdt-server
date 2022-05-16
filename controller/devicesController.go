package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iftdt/server/entity"
	"github.com/iftdt/server/service"
)

type DeviceController interface {
	Create(ctx *gin.Context) entity.Device
	Update(ctx *gin.Context)
	FindAll() []entity.Device
}

type deviceController struct {
	service service.DeviceService
}

func NewDeviceController(service service.DeviceService) DeviceController {
	return &deviceController{
		service: service,
	}
}

func (c *deviceController) FindAll() []entity.Device {
	return c.service.FindAll()
}

func (c *deviceController) Create(ctx *gin.Context) entity.Device {
	var device entity.Device
	ctx.BindJSON(&device)
	return c.service.Create(device)
}

func (c *deviceController) Update(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	device := c.service.FindOne(id)
	ctx.BindJSON(&device)
	c.service.Update(device)
}
