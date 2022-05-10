package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iftdt/server/common"
	"github.com/iftdt/server/controller"
)

func SetupRouters() *gin.Engine {
	gin.SetMode(common.ENV.GIN_MODE)
	app := gin.Default()
	api := app.Group("/api/v1")
	{
		// for admin
		api.GET("devices", controller.DeviceController{}.FindAll)
		// api.GET("notifications", controller.Notification{}.FindAll)
		// api.POST("notifications", controller.Notification{}.Create)

		// for app
		api.POST("devices", controller.DeviceController{}.Create)
		api.PUT("devices/:id", controller.DeviceController{}.Update)
	}
	return app
}
