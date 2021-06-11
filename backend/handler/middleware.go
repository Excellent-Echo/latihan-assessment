package handler

import (
	"latihanassesment/auth"
	"latihanassesment/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Middleware(userService user.Service, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || len(authHeader) == 0 {

			c.AbortWithStatusJSON(401, "error response")
			return
		}

		// eksekusi code untuk mengecek apakah token itu valid dari server kita atau tidak
		token, err := authService.ValidateToken(authHeader)

		if err != nil {

			c.AbortWithStatusJSON(401, "error response")
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.AbortWithStatusJSON(401, "error response")
			return
		}
		userID := int(claim["user_id"].(float64))

		//kita bisa pakai nanti

		c.Set("currentUser", userID)
	}
}
