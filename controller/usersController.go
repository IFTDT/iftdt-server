package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iftdt/server/entity"
	"github.com/iftdt/server/middleware"
	"github.com/iftdt/server/service"
)

type UserController interface {
	Create(ctx *gin.Context) entity.User
	FindAll() []entity.User
	Login(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) Create(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	return c.service.Create(user)
}

func (c *userController) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c *userController) Login(ctx *gin.Context) {
	var loginUser entity.User
	ctx.BindJSON(&loginUser)
	user, err := c.service.Login(loginUser)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	token, _ := middleware.SetToken(user.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"code":  200,
		"token": token,
	})
}
