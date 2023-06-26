package controller

import (
	"net/http"
	"server/model/web"
	"server/service"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	Service service.UserService
}

func NewHandler(s service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		Service: s,
	}
}

func (h *UserControllerImpl) CreateUser(c *gin.Context) {
	var u web.UserCreateRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)

}

func (h *UserControllerImpl) Login(c *gin.Context) {
	var user web.UserLoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.Service.Login(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return  
	}

	c.SetCookie("jwt", u.AccessToken, 3600, "/", "localhost", false, true)
	
	res := &web.UserLoginResponse{
		ID: u.ID,
		Username: u.Username,
	}
	c.JSON(http.StatusOK, res)
}

func (h *UserControllerImpl) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}