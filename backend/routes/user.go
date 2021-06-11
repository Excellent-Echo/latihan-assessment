package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marwanjuna/latihan-assessment/config"
	"github.com/marwanjuna/latihan-assessment/handler"
	"github.com/marwanjuna/latihan-assessment/user"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connect()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	userHandler             = handler.NewUserHandler(userService)
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", userHandler.ShowAllUsersHandler)
}
