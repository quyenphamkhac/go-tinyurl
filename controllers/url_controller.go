package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/quyenphamkhac/go-tinyurl/dtos"
	"github.com/quyenphamkhac/go-tinyurl/models"
	"github.com/quyenphamkhac/go-tinyurl/pkg/apperrors"
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

func (ctrl *URLController) CreateURL(c *gin.Context) {
	var createURLDto dtos.CreateURLDto
	if err := c.ShouldBindJSON(&createURLDto); err != nil {
		c.Error(apperrors.New(http.StatusBadRequest, err.Error()))
		return
	}
	userClaims, isExisted := c.Get("user")

	if !isExisted {
		c.Error(apperrors.New(http.StatusBadRequest, "User Claims Not Found"))
		return
	}
	user := userClaims.(*models.UserClaims)
	uuid, err := gocql.ParseUUID(user.UserID)
	if err != nil {
		c.Error(apperrors.New(http.StatusBadRequest, "User Claims Invalid"))
		return
	}
	url, err := ctrl.service.CreateURL(&createURLDto, &models.User{ID: uuid, Username: user.Username})
	if err != nil {
		c.Error(apperrors.New(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": url})
}

func (ctrl *URLController) GetURLByHash(c *gin.Context) {
	var getURLByHashDto dtos.GetURLByHashDto
	if err := c.ShouldBindUri(&getURLByHashDto); err != nil {
		c.Error(apperrors.New(http.StatusBadRequest, err.Error()))
		return
	}
	userClaims, isExisted := c.Get("user")

	if !isExisted {
		c.Error(apperrors.New(http.StatusBadRequest, "User Claims Not Found"))
		return
	}
	user := userClaims.(*models.UserClaims)
	url, err := ctrl.service.GetUserURLByHash(getURLByHashDto.Hash, user)
	if err != nil {
		c.Error(apperrors.New(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": url})
}

func (ctrl *URLController) RedirectURLByHash(c *gin.Context) {
	var getURLByHashDto dtos.GetURLByHashDto
	if err := c.ShouldBindUri(&getURLByHashDto); err != nil {
		c.Error(apperrors.New(http.StatusBadRequest, err.Error()))
		return
	}
	url, err := ctrl.service.GetURLByHash(getURLByHashDto.Hash)
	if err != nil {
		c.Error(apperrors.New(http.StatusInternalServerError, err.Error()))
		return
	}
	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}
