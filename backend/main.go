package main

import (
	"book-list/docs"
	"book-list/routes"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @securityDefinitions.apikey Auth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	docs.SwaggerInfo.Title = "book list API documentation"
	docs.SwaggerInfo.Description = "book list API documentation"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "book-list-helmi.herokuapp.com"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"https"}

	//routes API
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.UserRoute(r)
	routes.RouteBook(r)

	r.Run()
}
