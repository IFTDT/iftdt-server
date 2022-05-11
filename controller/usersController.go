package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iftdt/server/middleware"
	"github.com/iftdt/server/model"
)

type UserController struct{}

func (controller UserController) FindAll(c *gin.Context) {
	users, err := model.User{}.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": users,
		})
	}
}

func (controller UserController) Create(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := model.User{}.Create(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}

func (controller UserController) Update(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	user, err := model.User{}.FindOne(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(user)
	err = model.User{}.Update(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}

func (controller UserController) Login(c *gin.Context) {
	var loginUser model.User
	c.BindJSON(&loginUser)
	user, err := model.User{}.Login(loginUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, _ := middleware.SetToken(user.ID)
	c.JSON(http.StatusOK, gin.H{
		"data":  user,
		"token": token,
	})
}
