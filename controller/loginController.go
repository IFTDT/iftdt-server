package controller

import "github.com/gin-gonic/gin"

type LoginController interface {
	Login(ctx *gin.Context)
}

type loginController struct {
}

func NewLoginController() LoginController {
	return &loginController{}
}

func (c *loginController) Login(ctx *gin.Context) {

}
