package routes

import (
	"book-list/config"
	"book-list/handler"
	"book-list/layer/user"

	"github.com/gin-gonic/gin"
)

var (
	DB             = config.Connection()
	userRepository = user.NewRepository(DB)
	userService    = user.NewService(userRepository)
	userHandler    = handler.NewUserHandler(userService)
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", userHandler.ShowAllUser)
	r.POST("/users/register", userHandler.CreateNewUsers)
}
