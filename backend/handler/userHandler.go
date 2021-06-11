package handler

import (
	"github.com/gin-gonic/gin"
	"latihan-assessment/entity"
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
	users, err := h.userService.AllUser()

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

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var NewUserInput entity.UserInput

	if err := c.ShouldBindJSON(&NewUserInput); err != nil {
		splitError := helper.ErrorInformation(err)

		responseError := helper.APIResponse(
			"Input data required",
			400,
			"bad request",
			gin.H{
				"error": splitError,
			},
		)

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.userService.NewUser(NewUserInput)

	if err != nil {
		responseError := helper.APIResponse(
			"Internal server error",
			500,
			"error",
			gin.H{
				"error": err.Error(),
			},
		)

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse(
		"Success Create new Users Data",
		201,
		"Status Created",
		newUser,
	)

	c.JSON(201, response)
}

func (h *userHandler) UserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	userByID, err := h.userService.UserById(id)

	if err != nil {
		responseError := helper.APIResponse("Error bad request get user id", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("Success get user by id", 200, "success", userByID)
	c.JSON(200, response)
}

func (h *userHandler) DeleteUserIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	userDelete, err := h.userService.DeleteById(id)

	if err != nil {
		responseError := helper.APIResponse("Error delete user id", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	responseSuccess := helper.APIResponse("Success delete user id", 200, "Delete OK", userDelete)
	c.JSON(200, responseSuccess)
}

func (h *userHandler) UpdateUserIDHandler(c *gin.Context) {
	id := c.Params.ByName("user_id")

	var userUpdate entity.UserInputUpdate

	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		splitError := helper.ErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"error": splitError})

		c.JSON(400, responseError)
		return
	}

	users, err := h.userService.UpdateById(id, userUpdate)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update user by id", 200, "success", users)
	c.JSON(200, response)
}
