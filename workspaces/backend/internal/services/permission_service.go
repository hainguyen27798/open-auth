package services

import (
	"database/sql"
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/internal/db"
	"github.com/go-open-auth/internal/dto"
	"github.com/go-open-auth/internal/repos"
	"github.com/go-open-auth/pkg/response"
	"github.com/go-open-auth/pkg/utils"
)

type IPermissionService interface {
	CreateNewPermission(payload dto.PermissionRequestDTO) *int
	GetAllPermissions() []dto.PermissionResponseDTO
}

type permissionService struct {
	permissionRepo repos.IPermissionRepo
}

func NewPermissionService(permissionRepo repos.IPermissionRepo) IPermissionService {
	return &permissionService{
		permissionRepo,
	}
}

func (ps permissionService) CreateNewPermission(payload dto.PermissionRequestDTO) *int {
	err := ps.permissionRepo.CreateNewPermission(db.InsertNewPermissionParams{
		ServiceName: payload.ServiceName,
		Resource:    payload.Resource,
		Action:      payload.Action,
		Attributes:  payload.Attributes,
		Description: sql.NullString{
			String: payload.Description,
		},
	})
	if err != nil {
		global.Logger.Error(err.Error())
		return &[]int{response.ErrCreateFailed}[0]
	}
	return &[]int{response.CreatedSuccess}[0]
}

func (ps permissionService) GetAllPermissions() []dto.PermissionResponseDTO {
	return utils.ModelToDtos[dto.PermissionResponseDTO](ps.permissionRepo.GetAllPermission())
}
