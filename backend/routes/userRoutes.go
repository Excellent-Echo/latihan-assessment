package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"latihan-assessment/config"
)

var (
	DB *gorm.DB = config.ConnectToDatabase()
)

func UserRoute(route *gin.Engine)  {
	route.GET("/users")
}