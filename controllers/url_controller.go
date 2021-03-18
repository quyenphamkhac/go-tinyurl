package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/go-tinyurl/services"
)

type URLController struct {
	service *services.URLService
}

func NewURLController(s *services.URLService) *URLController {
	return &URLController{
		service: s,
	}
}

func (ctrl *URLController) GetAllURLs(c *gin.Context) {
	urls := ctrl.service.GetAllURLs()
	c.JSON(http.StatusOK, gin.H{"data": urls})
}

func (ctrl *URLController) HashURL(c *gin.Context) {

}
