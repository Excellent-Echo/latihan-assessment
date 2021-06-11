package main

import (
	"book-list/config"
	"book-list/migration"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB = config.Connect()

type User struct {
	gorm.Model
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date"`
}

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func main() {
	r := gin.Default()

	r.GET("/users", HandleGetUsers)
	r.POST("/user", HandlePostUser)
	r.Run(":4444")
}

func HandleGetUsers(c *gin.Context) {
	var users []migration.User

	if err := DB.Find(&users).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, users)
}

func HandlePostUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err.Error())
		return
	}

	if err := DB.Create(&user).Error; err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
