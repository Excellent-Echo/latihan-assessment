package handler

import (
	"book-list/config"
	"book-list/entity"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB = config.Connect()

func HandleGetUsers(c *gin.Context) {
	var users []entity.User

	if err := DB.Find(&users).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, users)
}

func HandleGetUserByID(c *gin.Context) {
	var user entity.User

	id := c.Params.ByName("id")

	if err := DB.Where("id = ?", id).Find(&user).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, user)
}

func HandlePostUser(c *gin.Context) {
	var user entity.User

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

func HandleDeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := DB.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "delete user successfully",
	})

}
