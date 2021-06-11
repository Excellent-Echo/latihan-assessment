package main

import (
	"book-list/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//routes API

	routes.UserRoute(r)

	r.Run(":5555")
}
