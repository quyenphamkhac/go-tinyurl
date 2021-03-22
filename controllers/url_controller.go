package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/go-tinyurl/dtos"
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

func (ctrl *URLController) CreateURL(c *gin.Context) {
	var createURLDto dtos.CreateURLDto
	if err := c.ShouldBindJSON(&createURLDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	url, err := ctrl.service.CreateURL(&createURLDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": url})
}
