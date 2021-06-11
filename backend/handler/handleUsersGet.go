package handler

import (
	"latihan-assessment/config"
	"latihan-assessment/entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB = config.ConnDB()

func HandleUsersGet(c *gin.Context) {
	var users []entities.Users

	if err := DB.Find(&users).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, users)
}