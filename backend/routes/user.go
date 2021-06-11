package routes

import (
	"backend/auth"
	"backend/config"
	"backend/handler"
	"backend/layer/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connection()
	authService             = auth.NewService()
	userService             = user.NewService(userRepository)
	userRepository          = user.NewRepository(DB)
	userHandler             = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", handler.Middleware(userService, authService), userHandler.ShowAllUserHandler)
	r.POST("/users/register", userHandler.RegisterHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
	r.GET("/users/:user_id", handler.Middleware(userService, authService), userHandler.GetUserByIDHandler)
	r.DELETE("/users/:user_id", handler.Middleware(userService, authService), userHandler.DeleteUserByIDHandler)
	r.PUT("/users/:user_id", handler.Middleware(userService, authService), userHandler.UpdateUserByIDHandler)
}
