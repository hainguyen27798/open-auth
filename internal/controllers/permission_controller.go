package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/open-auth/internal/dto"
	"github.com/open-auth/internal/services"
	"github.com/open-auth/pkg/response"
	"github.com/open-auth/pkg/utils"
)

type PermissionController struct {
	permissionService services.IPermissionService
}

func NewPermissionController(permissionService services.IPermissionService) *PermissionController {
	return &PermissionController{
		permissionService,
	}
}

func (pc *PermissionController) Create(c *gin.Context) {
	payload := utils.BodyToDto[dto.PermissionRequestDTO](c)
	if payload == nil {
		return
	}
	response.MessageResponse(c, pc.permissionService.CreateNewPermission(*payload).Code())
}

func (pc *PermissionController) Search(c *gin.Context) {
	query := utils.QueryToDto[dto.SearchDTO](c)
	permission := pc.permissionService.SearchPermissions(query)
	response.OkResponse(c, response.CodeSuccess, permission)
}

func (pc *PermissionController) GetAll(c *gin.Context) {
	permission := pc.permissionService.GetAllPermissions()
	response.OkResponse(c, response.CodeSuccess, permission)
}

func (pc *PermissionController) Update(c *gin.Context) {
	permissionId := c.Param("id")
	payload := utils.BodyToDto[dto.UpdatePermissionRequestDTO](c)
	if payload == nil {
		return
	}
	response.MessageResponse(c, pc.permissionService.UpdatePermission(permissionId, *payload).Code())
}

func (pc *PermissionController) Delete(c *gin.Context) {
	permissionId := c.Param("id")
	response.MessageResponse(c, pc.permissionService.DeletePermission(permissionId).Code())
}

func (pc *PermissionController) GetPermissionOptions(c *gin.Context) {
	roleId := c.Param("roleId")
	roles := pc.permissionService.GetPermissionOptions(roleId)
	response.OkResponse(c, response.CodeSuccess, roles)
}
