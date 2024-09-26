package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/open-auth/internal/dto"
	"github.com/open-auth/internal/services"
	"github.com/open-auth/pkg/response"
	"github.com/open-auth/pkg/utils"
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

	response.OkResponse(c, response.CodeSuccess, *user)
}

func (uc *UserController) GetUsers(c *gin.Context) {
	query := utils.QueryToDto[dto.SearchDTO](c)
	response.OkResponse(c, response.CodeSuccess, uc.userService.SearchUsers(query))
}
