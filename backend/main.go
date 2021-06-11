package main

import (
	"backend/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connnection()
	r := gin.Default()

	r.Run(":8000")
}
