package main

import (
	"github.com/gin-gonic/gin"

	"latihan-assessment/handler"
)


func main() {
	r := gin.Default()

	r.GET("/users", handler.HandleUsersGet )
	r.Run(":8080")

}
