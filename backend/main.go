package main

import (
	"books-project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.UserRoutes(r)
	r.Run(":8000")
}
