package handler

import (
	"book-list/auth"
	"book-list/entity"
	"book-list/layer/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service     user.Service
	authService auth.Service
}

func NewUserHandler(service user.Service, authService auth.Service) *userHandler {
	return &userHandler{service, authService}
}

func (h *userHandler) ShowAllUser(c *gin.Context) {
	users, err := h.service.GetAllUsers()

	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(200, users)
}

func (h *userHandler) CreateNewUsers(c *gin.Context) {
	var input entity.UserInput

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

func (h *userHandler) LoginUser(c *gin.Context) {
	var loginINput entity.LoginInput

	if err := c.ShouldBindJSON(&loginINput); err != nil {
		c.JSON(40, gin.H{"errors": err.Error()})
		return
	}

	userData, err := h.service.LoginUser(loginINput)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
	}

	token, err := h.authService.GenerateToken(userData.ID)

	var LoginOutput = user.UserLoginOutput{
		ID:            userData.ID,
		Name:          userData.Name,
		Address:       userData.Address,
		DateBirth:     user.FormatStringDateBirth(userData.DateBirth),
		Email:         userData.Email,
		Authorization: token,
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, LoginOutput)
}
