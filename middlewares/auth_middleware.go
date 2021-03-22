package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/go-tinyurl/services"
)

type authorizeHeader struct {
	Token string `header:"Authorization"`
}

func AuthorizeWithJwt(jwtService *services.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var h authorizeHeader
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		token := strings.Split(h.Token, "Bearer ")
		if len(token) < 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Must provide Authorization header with format `Bearer {token}`",
			})
			c.Abort()
			return
		}
		claims, err := jwtService.VerifyToken(token[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("user", claims.User)
	}
}
