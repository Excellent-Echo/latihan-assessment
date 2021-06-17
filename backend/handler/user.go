package handler

import (
	"book-list/auth"
	"book-list/entity"
	"book-list/layer/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service     user.Service
	authService auth.Service
}

func NewUserHandler(service user.Service, authService auth.Service) *userHandler {
	return &userHandler{service, authService}
}

// GetUsers godoc
// @Security Auth
// @Summary List existing users
// @Description Get all existing users
// @Tags users
// @Accept json
// @Produce json
// @success 200 {array} entity.UserOutput
// @Failure 500 {object} helper.Response
// @Router /users [get]
func (h *userHandler) ShowAllUser(c *gin.Context) {
	users, err := h.service.GetAllUsers()

	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}

	c.JSON(200, users)
}

// CreateUsers godoc
// @Security Auth
// @Summary Create new users
// @Description Create a new users
// @Tags users
// @Accept json
// @Produce json
// @Param user body entity.UserInput true "create user"
// @Success 200 {object} entity.UserOutput
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Router /users/register [post]
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

// Auth godoc
// @Summary Provides a JSON web Token
// @Description Authenticates a user and provide a JWT token
// @Tags users
// @ID Authentication
// @Consume json
// @Produce json
// @Param user body entity.LoginInput true "Login user"
// @Success 200 {object} user.UserLoginOutput
// @Failure 401 {object} helper.Response
// @Router /users/login [post]
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

// GetUserByID godoc
// @Security Auth
// @Summary Get user by ID
// @Description Get a user by user id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @success 200 {object} entity.UserOutput
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Router /users/{id} [get]
func (h *userHandler) ShowUserByID(c *gin.Context) {
	id := c.Param("user_id")

	user, err := h.service.GetUserByID(id)
	if err != nil {

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

// UpdateUserByID godoc
// @Security Auth
// @Summary Update user
// @Description Update a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "user ID"
// @Param user body entity.UserUpdateInput true "update user"
// @Success 200 {object} entity.UserOutput
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Router /users/{id} [put]
func (h *userHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Param("user_id")

	var updateUserInput entity.UserUpdateInput

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {

		c.JSON(400, gin.H{"error": err})
		return
	}

	idParam, _ := strconv.Atoi(id)

	// authorization userid dari params harus sama dengan user id yang login
	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {

		c.JSON(401, gin.H{"error": "unauthorize"})
		return
	}

	user, err := h.service.UpdateUserByID(id, updateUserInput)
	if err != nil {

		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, user)
}

// DeleteUser godoc
// @Security Auth
// @Summary Delete user by id
// @Description Delete a user by id
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "user ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Router /users/{id} [delete]
func (h *userHandler) DeleterUserByID(c *gin.Context) {
	id := c.Param("user_id")

	user, err := h.service.DeleteUserByID(id)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}
