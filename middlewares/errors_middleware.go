package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/go-tinyurl/pkg/apperrors"
)

func ErrorsMiddleware(errorType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		var appError *apperrors.AppError
		errors := c.Errors.ByType(errorType)
		if len(errors) > 0 {
			err := errors[0].Err
			switch err := err.(type) {
			case *apperrors.AppError:
				appError = err
			default:
				appError = &apperrors.AppError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}
			c.IndentedJSON(appError.Code, appError)
			c.Abort()
			return
		}
	}
}
