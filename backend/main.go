package main

import (
	"github.com/gin-gonic/gin"
	"latihan-assessment/routes"
)

func main() {
	router := gin.Default()
	routes.UserRoute(router)

	router.Run()
}
