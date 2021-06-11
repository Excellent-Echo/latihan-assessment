package main

import (
	"github.com/gin-gonic/gin"

	"latihan-assessment/handler"
)


func main() {
	r := gin.Default()

	r.GET("/users", handler.HandleUsersGet )
	r.GET("/users/:user_id", handler.HandleUsersIdGet )

	r.Run(":8080")

}
