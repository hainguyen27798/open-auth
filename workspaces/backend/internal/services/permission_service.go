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
	CreateNewPermission(payload dto.PermissionRequestDTO) *response.ServerCode
	GetAllPermissions() []dto.PermissionResponseDTO
	UpdatePermission(id string, payload dto.UpdatePermissionRequestDTO) *response.ServerCode
	DeletePermission(id string) *response.ServerCode
}

type permissionService struct {
	permissionRepo repos.IPermissionRepo
}

func NewPermissionService(permissionRepo repos.IPermissionRepo) IPermissionService {
	return &permissionService{
		permissionRepo,
	}
}

func (ps *permissionService) CreateNewPermission(payload dto.PermissionRequestDTO) *response.ServerCode {
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
		return response.ReturnCode(response.ErrCreateFailed)
	}
	return response.ReturnCode(response.CreatedSuccess)
}

func (ps *permissionService) GetAllPermissions() []dto.PermissionResponseDTO {
	return utils.ModelToDtos[dto.PermissionResponseDTO](ps.permissionRepo.GetAllPermission())
}

func (ps *permissionService) UpdatePermission(id string, payload dto.UpdatePermissionRequestDTO) *response.ServerCode {
	updatePayloadDto, errCode := utils.DtoToModel[db.UpdatePermissionParams](payload)
	updatePayloadDto.ID = id

	if errCode != nil {
		return errCode
	}

	ok, err := ps.permissionRepo.UpdatePermission(*updatePayloadDto)
	if err != nil {
		global.Logger.Error(err.Error())
		return response.ReturnCode(response.ErrBadRequest)
	}

	if !ok {
		return response.ReturnCode(response.ErrNotFound)
	}

	return response.ReturnCode(response.CodeSuccess)
}

func (ps *permissionService) DeletePermission(id string) *response.ServerCode {
	if ok := ps.permissionRepo.DeletePermission(id); ok {
		return response.ReturnCode(response.CodeSuccess)
	}
	return response.ReturnCode(response.ErrNotFound)
}
