package handler

import (
	"book-list/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Middleware(authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		Header := c.GetHeader("Authorization")

		if Header == "" || len(Header) == 0 {
			c.AbortWithStatusJSON(401, gin.H{"errors": "unauhorize"})
			return
		}

		token, err := authService.ValidateToken(Header)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "unauthorize"})
			return
		}

		userID := int(claim["user_id"].(float64))

		c.Set("currentUser", userID)
	}
}
