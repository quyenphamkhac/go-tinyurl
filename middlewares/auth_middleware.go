package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/go-tinyurl/services"
)

func AuthorizeWithJwt(jwtService *services.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Must provide Authorization header with format `Bearer {token}`",
			})
			return
		}
		token := authHeader[len("Bearer")+1:]
		claims, err := jwtService.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Set("user", claims.User)
		c.Next()
	}
}
