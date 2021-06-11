package main

import (
	"backend/handler"
	"backend/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connect()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	userHandler             = handler.NewUserHandler(userService)
)

func main() {

	r := gin.Default()
	r.GET("/user", userHandler, ShowUser)

	r.Run("8080")

}
