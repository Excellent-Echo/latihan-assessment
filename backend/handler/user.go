package handler

import (
	"latihanassesment/auth"
	"latihanassesment/entity"
	"latihanassesment/helper"
	"latihanassesment/user"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) ShowUserHandler(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {

		c.JSON(500, users)
		return
	}

	c.JSON(200, users)
}

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput
	var birth time.Time

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		splitError := helper.SpilitErrorInformation(err)

		c.JSON(400, splitError)
		return
	}

	newUser, err := h.userService.SaveNewUser(birth, inputUser)
	if err != nil {
		c.JSON(500, newUser)
		return
	}

	c.JSON(201, newUser)
}

func (h *userHandler) GetUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	user, err := h.userService.GetUserByID(id)
	if err != nil {

		c.JSON(400, user)
		return
	}

	c.JSON(200, user)
}

func (h *userHandler) DeleteUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	user, err := h.userService.DeleteUserByID(id)

	if err != nil {
		c.JSON(400, user)
		return
	}

	c.JSON(200, user)
}

func (h *userHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	var updateUserInput entity.UpdateUserInput

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		c.JSON(400, updateUserInput)
		return
	}

	idParam, _ := strconv.Atoi(id)

	// authorization userid dari params harus sama dengan user id yang login
	UserData := int(c.MustGet("currentUser").(int))

	if idParam != UserData {

		c.JSON(401, updateUserInput)
		return
	}

	user, err := h.userService.UpdateUserByID(id, updateUserInput)
	if err != nil {

		c.JSON(500, user)
		return
	}
	c.JSON(200, user)
}

func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var LoginUser entity.LoginUserInput

	if err := c.ShouldBindJSON(&LoginUser); err != nil {

		c.JSON(400, LoginUser)
		return
	}

	userData, err := h.userService.LoginUser(LoginUser)

	if err != nil {

		c.JSON(401, LoginUser)
		return
	}

	token, err := h.authService.GenerateToken(userData.ID)
	if err != nil {

		c.JSON(401, token)
		return
	}
	c.JSON(200, token)
}
