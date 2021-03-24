package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/go-tinyurl/dtos"
	"github.com/quyenphamkhac/go-tinyurl/pkg/apperrors"
	"github.com/quyenphamkhac/go-tinyurl/services"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(s *services.AuthService) *AuthController {
	return &AuthController{
		service: s,
	}
}

func (ctrl *AuthController) SignUp(c *gin.Context) {
	var userDto dtos.SignUpDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.Error(apperrors.New(http.StatusBadRequest, err.Error()))
		return
	}
	user, err := ctrl.service.SignUp(&userDto)
	if err != nil {
		c.Error(apperrors.New(http.StatusBadRequest, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var credentials dtos.SignInDto
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.Error(apperrors.New(http.StatusBadRequest, err.Error()))
		return
	}
	accessTokenResp, err := ctrl.service.Login(&credentials)
	if err != nil {
		c.Error(apperrors.New(http.StatusUnauthorized, err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": accessTokenResp})
}
