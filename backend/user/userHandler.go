package user

import (
	"latihan-assessment/backend/user"
	"latihan-assessment/backend/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userhandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userhandler {
	return &userhandler(userService)
}

func (h *userhandler) ShowAllUser(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error in internal server"
		})
		return
	}

	UserResponse := helper.APIResponse("success get all users", 200, "success", users)
	c.JSON(http.StatusOK, users)
}
