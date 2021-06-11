package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"latihan-assessment/backend/config"
)

var DB *gorm.DB = config.Conn()

func main() {
	r := gin.Default()

	r.GET("/users", HandleUsers)
	r.POST("/users/register")
	r.POST("/users/login")
	r.GET("/users/:id")
	r.PUT("/users/:id")
	r.DELETE("/users/:id")

	r.GET("/books", HandleBook)
	r.POST("/books/")
	r.GET("/books/:id")
	r.PUT("/books/:id")
	r.DELETE("/books/:id")

	r.Run(":8080")
}

func HandleUsers(c *gin.Context) {
	var users []Users

	if err := DB.Find(&users).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, users)
}

func HandleBook(c *gin.Context) {

}
