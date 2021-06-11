package main

import (
	"book-list/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func main() {
	r := gin.Default()

	r.GET("/users", handler.HandleGetUsers)
	r.GET("/user/:id", handler.HandleGetUserByID)
	r.POST("/user", handler.HandlePostUser)
	r.DELETE("/user/:id", handler.HandleDeleteUser)
	r.Run(":4444")
}
