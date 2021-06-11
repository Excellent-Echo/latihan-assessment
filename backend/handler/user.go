package handler

import (
	"book-list/layer/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) ShowAllUser(c *gin.Context) {
	users, err := h.service.GetAllUsers()

	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(200, users)
}

func (h *userHandler) CreateNewUsers(c *gin.Context) {
	var input user.UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newUser, err := h.service.CreateNewUser(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, newUser)

}
