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
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token := authHeader[len("Bearer")+1:]
		jwtToken, err := jwtService.VerifyToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if !jwtToken.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
