package main

import (
	"os"

	"github.com/marwanjuna/latihan-assessment/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.UserRoute(r)
	routes.BookRoute(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
