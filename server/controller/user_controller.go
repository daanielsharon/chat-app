package controller

import "github.com/gin-gonic/gin"

type UserController interface{
	CreateUser(c *gin.Context)
	Login(c *gin.Context) 
	Logout(c *gin.Context)
}