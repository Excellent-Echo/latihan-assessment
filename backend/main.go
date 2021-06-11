package main

import (
	"github.com/marwanjuna/latihan-assessment/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.UserRoute(r)
	routes.BookRoute(r)

	r.Run(":4444")
}
