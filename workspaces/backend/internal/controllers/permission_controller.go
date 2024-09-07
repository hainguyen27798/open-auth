package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-open-auth/internal/dto"
	"github.com/go-open-auth/internal/services"
	"github.com/go-open-auth/pkg/response"
	"github.com/go-open-auth/pkg/utils"
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
	response.MessageResponse(c, *ps.permissionService.CreateNewPermission(*payload))
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
	response.MessageResponse(c, *ps.permissionService.UpdatePermission(permissionId, *payload))
}

func (ps *PermissionController) Delete(c *gin.Context) {
	permissionId := c.Param("id")
	response.MessageResponse(c, *ps.permissionService.DeletePermission(permissionId))
}
