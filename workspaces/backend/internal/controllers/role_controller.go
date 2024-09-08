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
	response.MessageResponse(c, rc.roleService.CreateNewRole(*payload).Code())
}

func (rc *RoleController) GetAll(c *gin.Context) {
	roles := rc.roleService.GetAllRoles()
	response.OkResponse(c, response.CodeSuccess, roles)
}

func (rc *RoleController) Get(c *gin.Context) {
	roleId := c.Param("id")
	role, errCode := rc.roleService.GetRole(roleId)

	if errCode != nil {
		response.NotFoundException(c, errCode.Code())
		return
	}

	response.OkResponse(c, response.CodeSuccess, role)
}
