package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iftdt/server/common"
	"github.com/iftdt/server/controller"
	"github.com/iftdt/server/dto"
	"github.com/iftdt/server/middleware"
	"github.com/iftdt/server/repository"
	"github.com/iftdt/server/service"
)

var (
	deviceRepository repository.DeviceRepository = repository.NewDeviceRepository(dto.DB)
	deviceService    service.DeviceService       = service.NewDeviceService(deviceRepository)
	deviceController controller.DeviceController = controller.NewDeviceController(deviceService)
)

var (
	userRepository repository.UserRepository = repository.NewUserRepository(dto.DB)
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)
)

var (
	notificationRepository repository.NotificationRepository = repository.NewNotificationRepository(dto.DB)
	notificationService    service.NotificationService       = service.NewNotificationService(notificationRepository, deviceRepository)
	notificationController controller.NotificationController = controller.NewNotificationController(notificationService)
)

func SetupRouters() *gin.Engine {
	gin.SetMode(common.ENV.GIN_MODE)
	app := gin.Default()

	app.Use(middleware.CORSMiddleware())

	auth := app.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{
		// for admin
		auth.GET("devices", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"data": deviceController.FindAll(),
			})
		})

		auth.GET("users", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"data": userController.FindAll(),
			})
		})

		auth.GET("notifications", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"data": notificationController.FindAll(),
			})
		})
		auth.POST("notifications", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, notificationController.Create(ctx))
		})

		// for app
		auth.POST("devices", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, deviceController.Create(ctx))
		})
		auth.PUT("devices/:id", func(ctx *gin.Context) {
			deviceController.Update(ctx)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		})
	}
	api := app.Group("/api/v1")
	{
		api.POST("login", userController.Login)
		api.POST("users", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, userController.Create(ctx))
		})
	}

	return app
}
