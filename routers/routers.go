package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iftdt/server/common"
	"github.com/iftdt/server/controller"
	"github.com/iftdt/server/middleware"
)

func SetupRouters() *gin.Engine {
	gin.SetMode(common.ENV.GIN_MODE)
	app := gin.Default()

	app.Use(middleware.CORSMiddleware())

	auth := app.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{
		// for admin
		auth.GET("devices", controller.DeviceController{}.FindAll)

		auth.GET("users", controller.UserController{}.FindAll)

		auth.GET("notifications", controller.NotificationController{}.FindAll)
		auth.POST("notifications", controller.NotificationController{}.Create)

		// for app
		auth.PUT("users/:id", controller.UserController{}.Update)

		auth.POST("devices", controller.DeviceController{}.Create)
		auth.PUT("devices/:id", controller.DeviceController{}.Update)
	}
	api := app.Group("/api/v1")
	{
		api.POST("login", controller.UserController{}.Login)
		api.POST("users", controller.UserController{}.Create)
	}

	return app
}
