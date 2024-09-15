package services

import (
	"github.com/open-auth/global"
	"github.com/open-auth/internal/dto"
	"github.com/open-auth/internal/models"
	"github.com/open-auth/internal/repos"
	"github.com/open-auth/pkg/response"
	"github.com/open-auth/pkg/utils"
)

type IPermissionService interface {
	CreateNewPermission(payload dto.PermissionRequestDTO) *response.ServerCode
	GetAllPermissions(payload dto.SearchDTO) dto.PaginationDto[dto.PermissionResponseDTO]
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
	newPayload, errCode := utils.DtoToModel[models.InsertNewPermissionParams](payload)

	if errCode != nil {
		return errCode
	}

	if err := ps.permissionRepo.CreateNewPermission(*newPayload); err != nil {
		global.Logger.Error(err.Error())
		return response.ReturnCode(response.ErrCreateFailed)
	}

	return response.ReturnCode(response.CreatedSuccess)
}

func (ps *permissionService) GetAllPermissions(payload dto.SearchDTO) dto.PaginationDto[dto.PermissionResponseDTO] {
	permission, total := ps.permissionRepo.GetAllPermission(
		payload.Search,
		payload.By,
		payload.Skip(),
		payload.Limit(),
	)
	return utils.ModelToPaginationDto[dto.PermissionResponseDTO](
		permission,
		dto.PaginationMetaDataDto{
			Total: total,
		},
	)
}

func (ps *permissionService) UpdatePermission(id string, payload dto.UpdatePermissionRequestDTO) *response.ServerCode {
	updatePayload, errCode := utils.DtoToModel[models.UpdatePermissionParams](payload)
	if errCode != nil {
		return errCode
	}

	updatePayload.ID = &id
	ok, err := ps.permissionRepo.UpdatePermission(*updatePayload)
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
