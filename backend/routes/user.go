package routes

import (
	"book-list/auth"
	"book-list/config"
	"book-list/handler"
	"book-list/layer/user"

	"github.com/gin-gonic/gin"
)

var (
	DB             = config.Connection()
	userRepository = user.NewRepository(DB)
	userService    = user.NewService(userRepository)
	authService    = auth.NewService()
	userHandler    = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", handler.Middleware(authService), userHandler.ShowAllUser)
	r.GET("/users/:user_id", handler.Middleware(authService), userHandler.ShowUserByID)
	r.POST("/users/register", userHandler.CreateNewUsers)
	r.POST("/users/login", userHandler.LoginUser)
	r.PUT("/users/:user_id", handler.Middleware(authService), userHandler.UpdateUserByIDHandler)
	r.DELETE("users/:user_id", handler.Middleware(authService), userHandler.DeleterUserByID)
}
