package routes

import (
	"book-list/config"
	"book-list/handler"
	"book-list/layer/user"
)

var (
	DB             = config.Connection()
	userRepository = user.NewRepository(DB)
	userService    = user.NewService(userRepository)
	userHandler    = handler.NewUserHandler(userService)
)
