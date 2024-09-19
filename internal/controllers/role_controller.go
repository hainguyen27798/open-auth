package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/open-auth/internal/dto"
	"github.com/open-auth/internal/services"
	"github.com/open-auth/pkg/response"
	"github.com/open-auth/pkg/utils"
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
	query := utils.QueryToDto[dto.SearchDTO](c)
	roles := rc.roleService.GetAllRoles(query)
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

func (rc *RoleController) GetRolePermissions(c *gin.Context) {
	roleId := c.Param("id")
	roles := rc.roleService.GetRolePermissions(roleId)
	response.OkResponse(c, response.CodeSuccess, roles)
}

func (rc *RoleController) Update(c *gin.Context) {
	roleId := c.Param("id")
	payload := utils.BodyToDto[dto.UpdateRoleRequestDTO](c)
	if payload == nil {
		return
	}
	response.MessageResponse(c, rc.roleService.UpdateRole(roleId, *payload).Code())
}

func (rc *RoleController) Delete(c *gin.Context) {
	roleId := c.Param("id")
	response.MessageResponse(c, rc.roleService.DeleteRole(roleId).Code())
}

func (rc *RoleController) AddRolePermission(c *gin.Context) {
	payload := utils.BodyToDto[dto.AddRolePermissionRequestDTO](c)
	roleId := c.Param("id")
	if payload == nil {
		return
	}

	resCode := rc.roleService.AddRolePermission(roleId, payload.PermissionId)

	response.MessageResponse(c, resCode.Code())
}
