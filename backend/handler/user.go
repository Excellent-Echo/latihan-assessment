package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marwanjuna/latihan-assessment/user"
)

type userHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) ShowAllUsersHandler(c *gin.Context) {
	users, err := h.userService.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
