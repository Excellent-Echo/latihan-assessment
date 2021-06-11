package main

import "github.com/gin-gonic/gin
		backend/config
"

func main() {
	config.Connnection()
	r := gin.Default()

	r.Run(":1000")
}
