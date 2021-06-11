package handler

import (
	"github.com/gin-gonic/gin"
	"latihan-assessment/helper"
	"latihan-assessment/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHander(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) ShowAllUserHandler(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error internal server",
			"error":   err.Error(),
		})
		return
	}

	response := helper.APIResponse(
		"Success Get All Users Data",
		200,
		"Status OK",
		users,
	)

	c.JSON(200, response)
}
