package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/dto"
	"github.com/open-auth/internal/services"
	"github.com/open-auth/pkg/response"
	"github.com/open-auth/pkg/utils"
)

type AuthController struct {
	authService services.IAuthService
}

func NewAuthController(authService services.IAuthService) *AuthController {
	return &AuthController{
		authService,
	}
}

func (ac *AuthController) Register(c *gin.Context) {
	params := utils.BodyToDto[dto.UserRegistrationRequestDTO](c)

	if params == nil {
		return
	}

	response.MessageResponse(c, ac.authService.Register(*params).Code())
}

func (ac *AuthController) Login(c *gin.Context) {
	params := utils.BodyToDto[dto.UserLoginRequestDTO](c)

	if params == nil {
		return
	}

	if res, errCode := ac.authService.Login(*params); errCode != nil {
		response.MessageResponse(c, errCode.Code())
	} else {
		response.OkResponse(c, response.LoginSuccess, *res)
	}
}

func (ac *AuthController) RefreshToken(c *gin.Context) {
	refreshToken := c.GetHeader(global.RefreshTokenKey)

	newToken, errCode := ac.authService.RefreshToken(refreshToken)
	if errCode != nil {
		response.MessageResponse(c, errCode.Code())
	} else {
		response.OkResponse(c, response.CodeSuccess, newToken)
	}
}

func (ac *AuthController) Logout(c *gin.Context) {
	refreshToken := c.GetHeader(global.RefreshTokenKey)
	code := ac.authService.Logout(refreshToken)
	response.MessageResponse(c, code.Code())
}
