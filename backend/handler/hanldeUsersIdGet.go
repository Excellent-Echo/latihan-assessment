package handler

import (
	"latihan-assessment/entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func HandleUsersIdGet(c *gin.Context) {
	var users entities.Users

	id := c.Params.ByName("user_id")

	if err := DB.Where("id = ?", id).Find(&users).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, users)
}