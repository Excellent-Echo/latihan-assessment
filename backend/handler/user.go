package handler

import (
	"backend/entity"
	"backend/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	response, err := h.userService.SaveNewUser(inputUser)
	if err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(201, response)
}
