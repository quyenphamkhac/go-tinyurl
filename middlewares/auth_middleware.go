package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/go-tinyurl/pkg/apperrors"
	"github.com/quyenphamkhac/go-tinyurl/services"
)

type authorizeHeader struct {
	Token string `header:"Authorization"`
}

func AuthorizeWithJwt(jwtService *services.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var h authorizeHeader
		if err := c.ShouldBindHeader(&h); err != nil {
			c.Error(apperrors.New(http.StatusBadRequest, err.Error()))
			return
		}
		token := strings.Split(h.Token, "Bearer ")
		if len(token) < 2 {
			c.Error(apperrors.New(http.StatusUnauthorized, "Must provide Authorization header with format `Bearer {token}`"))
			return
		}
		claims, err := jwtService.VerifyToken(token[1])
		if err != nil {
			c.Error(apperrors.New(http.StatusUnauthorized, err.Error()))
			return
		}
		c.Set("user", claims.User)
	}
}
