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
