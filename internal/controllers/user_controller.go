package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-open-auth/internal/dto"
	"github.com/go-open-auth/internal/services"
	"github.com/go-open-auth/pkg/response"
	"github.com/go-open-auth/pkg/utils"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService,
	}
}

func (uc *UserController) GetMe(c *gin.Context) {
	userEmail := c.GetHeader("userEmail")
	user, errCode := uc.userService.GetMe(userEmail)

	if errCode != nil {
		response.NotFoundException(c, errCode.Code())
		return
	}

	response.OkResponse(c, response.CodeSuccess, utils.ModelToDto[dto.UserResponseDTO](*user))
}
