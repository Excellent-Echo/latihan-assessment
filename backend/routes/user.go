package routes

import (
	"books-project/config"
	"books-project/handler"
	"books-project/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connection()
	userRepository          = user.NewRepository(DB)
	userService             = user.UserNewService(userRepository)
	userHandler             = handler.NewUserHandler(userService)
)

func UserRoutes(r *gin.Engine) {
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.GET("/users", userHandler.ShowAllUserHandler)
}
