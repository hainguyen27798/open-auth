package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-open-auth/internal/dto"
	"github.com/go-open-auth/internal/services"
	"github.com/go-open-auth/pkg/response"
	"github.com/go-open-auth/pkg/utils"
)

type RoleController struct {
	roleService services.IRoleService
}

func NewRoleController(roleService services.IRoleService) *RoleController {
	return &RoleController{
		roleService,
	}
}

func (rc *RoleController) Create(c *gin.Context) {
	payload := utils.BodyToDto[dto.RoleRequestDTO](c)
	if payload == nil {
		return
	}
	response.MessageResponse(c, *rc.roleService.CreateNewRole(*payload))
}

func (rc *RoleController) GetAll(c *gin.Context) {
	roles := rc.roleService.GetAllRole()
	response.OkResponse(c, response.CodeSuccess, roles)
}
