package main

import (
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ini routingan
	routes.UserRoute(r)

	r.Run()
}
