package helper

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"latihan-assessment/backend/config"
	"latihan-assessment/backend/entity"
)

var DB *gorm.DB = config.Conn()

func HandleUsers(c *gin.Context) {
	var users []entity.Users

	if err := DB.Find(&users).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, users)
}

func HandleUserID(c *gin.Context) {
	var users []entity.Users

	id := c.Params.ByName("id")

	if err := DB.Where("id = ?", id).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, users)
}

func HandleUsersPOST(c *gin.Context) {
	var users []entity.Users

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err.Error())
		return
	}

	if err := DB.Create(&user).Error; err != nil {
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func HandleBook(c *gin.Context) {

}
