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

func (ps *PermissionController) Create(c *gin.Context) {
	payload := utils.BodyToDto[dto.PermissionRequestDTO](c)
	if payload == nil {
		return
	}
	response.MessageResponse(c, ps.permissionService.CreateNewPermission(*payload).Code())
}

func (ps *PermissionController) Search(c *gin.Context) {
	query := utils.QueryToDto[dto.SearchDTO](c)
	permission := ps.permissionService.SearchPermissions(query)
	response.OkResponse(c, response.CodeSuccess, permission)
}

func (ps *PermissionController) GetAll(c *gin.Context) {
	permission := ps.permissionService.GetAllPermissions()
	response.OkResponse(c, response.CodeSuccess, permission)
}

func (ps *PermissionController) Update(c *gin.Context) {
	permissionId := c.Param("id")
	payload := utils.BodyToDto[dto.UpdatePermissionRequestDTO](c)
	if payload == nil {
		return
	}
	response.MessageResponse(c, ps.permissionService.UpdatePermission(permissionId, *payload).Code())
}

func (ps *PermissionController) Delete(c *gin.Context) {
	permissionId := c.Param("id")
	response.MessageResponse(c, ps.permissionService.DeletePermission(permissionId).Code())
}
