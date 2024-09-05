package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-open-auth/internal/dto"
	"github.com/go-open-auth/internal/services"
	"github.com/go-open-auth/pkg/response"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params dto.UserRegistrationRequestDTO
	if err := c.ShouldBindBodyWithJSON(&params); err != nil {
		response.ValidateErrorResponse(c, err)
		return
	}

	response.MessageResponse(c, uc.userService.Register(params))
}

func (uc *UserController) Login(c *gin.Context) {
	var params dto.UserRegistrationRequestDTO
	if err := c.ShouldBindBodyWithJSON(&params); err != nil {
		response.ValidateErrorResponse(c, err)
		return
	}

	response.MessageResponse(c, uc.userService.Register(params))
}
