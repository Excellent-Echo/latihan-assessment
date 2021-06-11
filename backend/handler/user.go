package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marwanjuna/latihan-assessment/auth"
	"github.com/marwanjuna/latihan-assessment/entity"
	"github.com/marwanjuna/latihan-assessment/helper"
	"github.com/marwanjuna/latihan-assessment/user"
)

type userHandler struct {
	userService user.UserService
	authService auth.Service
}

func NewUserHandler(userService user.UserService, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": splitError,
		})
		return
	}

	response, err := h.userService.CreateNewUser(inputUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, response)
}

func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var inputLoginUser entity.LoginUserInput

	if err := c.ShouldBindJSON(&inputLoginUser); err != nil {
		splitError := helper.SplitErrorInformation(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": splitError,
		})
		return
	}

	userData, err := h.userService.LoginUser(inputLoginUser)

	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := h.authService.GenerateToken(int(userData.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":            userData.ID,
		"name":          userData.Name,
		"address":       userData.Address,
		"date_birth":    userData.DateBirth,
		"email":         userData.Email,
		"authorization": token,
	})
}

func (h *userHandler) ShowUserByIDhandler(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Param("id")

	var updateUserInput entity.UpdateUserInput

	idParam, _ := strconv.Atoi(id)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		c.JSON(401, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	user, err := h.userService.UpdateUserByID(id, updateUserInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) DeleteByUserIDHandler(c *gin.Context) {
	id := c.Param("id")

	idParam, _ := strconv.Atoi(id)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		c.JSON(401, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	user, err := h.userService.DeleteUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
