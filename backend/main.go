package main

import (
	"latihan-assessment/backend/config"
	"latihan-assessment/backend/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB = config.Conn()
)

func main() {
	r := gin.Default()

	r.GET("/users", helper.HandleUsers)
	r.POST("/users/register", helper.HandleUsersPOST)
	r.POST("/users/login")
	r.GET("/users/:id", helper.HandleUserID)
	r.PUT("/users/:id", helper.HandleUserUpdate)
	r.DELETE("/users/:id", helper.HandleUserDelete)

	r.GET("/books", helper.HandleBook)
	r.POST("/books/")
	r.GET("/books/:id")
	r.PUT("/books/:id")
	r.DELETE("/books/:id")

	r.Run(":8080")
}
