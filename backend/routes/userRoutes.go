package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"latihan-assessment/config"
	"latihan-assessment/handler"
	"latihan-assessment/user"
)

var (
	DB             *gorm.DB = config.ConnectToDatabase()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	userHandler             = handler.NewUserHander(userService)
)

func UserRoute(route *gin.Engine) {
	route.GET("/users", userHandler.ShowAllUserHandler)
	route.POST("/users/register", userHandler.CreateUserHandler)
	//route.POST("/users/login", userHandler.ShowAllUserHandler)
	route.GET("users/:user_id", userHandler.UserByIDHandler)
	route.PUT("users/:user_id", userHandler.UpdateUserIDHandler)
	route.DELETE("users/:user_id", userHandler.DeleteUserIDHandler)
}
