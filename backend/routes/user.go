package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marwanjuna/latihan-assessment/auth"
	"github.com/marwanjuna/latihan-assessment/config"
	"github.com/marwanjuna/latihan-assessment/handler"
	"github.com/marwanjuna/latihan-assessment/user"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connect()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	authService             = auth.NewService()
	userHandler             = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
	r.GET("/users", userHandler.ShowAllUsersHandler)
	r.GET("/users/:id", userHandler.ShowUserByIDhandler)
	r.PUT("/users/:id", userHandler.UpdateUserByIDHandler)
	r.DELETE("/users/:id", userHandler.DeleteByUserIDHandler)

}
