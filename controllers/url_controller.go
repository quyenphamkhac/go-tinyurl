package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/quyenphamkhac/go-tinyurl/dtos"
	"github.com/quyenphamkhac/go-tinyurl/entities"
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
	userClaims, isExisted := c.Get("user")

	if !isExisted {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't get context user"})
		return
	}
	user, ok := userClaims.(*entities.UserClaims)
	fmt.Println(userClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "context user invalid"})
		return
	}
	uuid, err := gocql.ParseUUID(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant parse uuid"})
		return
	}
	url, err := ctrl.service.CreateURL(&createURLDto, &entities.User{ID: uuid, Username: user.Username})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": url})
}

func (ctrl *URLController) GetURLByHash(c *gin.Context) {
	var getURLByHashDto dtos.GetURLByHashDto
	if err := c.ShouldBindUri(&getURLByHashDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userClaims, isExisted := c.Get("user")

	if !isExisted {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't get context user"})
		return
	}
	user := userClaims.(*entities.UserClaims)
	url, err := ctrl.service.GetURLByHash(getURLByHashDto.Hash, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": url})
}
