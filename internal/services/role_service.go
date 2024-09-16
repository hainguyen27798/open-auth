package services

import (
	"github.com/open-auth/global"
	"github.com/open-auth/internal/dto"
	"github.com/open-auth/internal/models"
	"github.com/open-auth/internal/repos"
	"github.com/open-auth/pkg/response"
	"github.com/open-auth/pkg/utils"
	"go.uber.org/zap"
)

type IRoleService interface {
	CreateNewRole(payload dto.RoleRequestDTO) *response.ServerCode
	GetAllRoles() []dto.RoleResponseDTO
	GetRole(id string) (*dto.RoleResponseDTO, *response.ServerCode)
	DeleteRole(id string) *response.ServerCode
	UpdateRole(id string, payload dto.UpdateRoleRequestDTO) *response.ServerCode
}

type roleService struct {
	roleRepo repos.IRoleRepo
}

func NewRoleService(roleRepo repos.IRoleRepo) IRoleService {
	return &roleService{
		roleRepo,
	}
}

func (rs *roleService) CreateNewRole(payload dto.RoleRequestDTO) *response.ServerCode {
	payloadRequest, errCode := utils.DtoToModel[models.InsertNewRoleParams](payload)

	if errCode != nil {
		return errCode
	}

	if err := rs.roleRepo.CreateNewRole(*payloadRequest); err != nil {
		global.Logger.Error("CreateNewRole: ", zap.Error(err))
		return response.ReturnCode(response.ErrCreateFailed)
	}

	return response.ReturnCode(response.CreatedSuccess)
}

func (rs *roleService) GetAllRoles() []dto.RoleResponseDTO {
	return utils.ModelToDtos[dto.RoleResponseDTO](rs.roleRepo.GetAllRoles())
}

func (rs *roleService) GetRole(id string) (*dto.RoleResponseDTO, *response.ServerCode) {
	role, err := rs.roleRepo.GetById(id)

	if err != nil {
		return nil, response.ReturnCode(response.ErrNotFound)
	}

	return utils.ModelToDto[dto.RoleResponseDTO](*role), nil
}

func (rs *roleService) UpdateRole(id string, payload dto.UpdateRoleRequestDTO) *response.ServerCode {
	updatePayload, err := utils.DtoToModel[models.UpdateRoleParams](payload)
	updatePayload.ID = &id
	if err != nil {
		return response.ReturnCode(response.ErrCodeParamInvalid)
	}

	if _, err := rs.roleRepo.Update(*updatePayload); err != nil {
		return response.ReturnCode(response.ErrBadRequest)
	}

	return response.ReturnCode(response.CodeSuccess)
}

func (rs *roleService) DeleteRole(id string) *response.ServerCode {
	ok := rs.roleRepo.Delete(id)

	if !ok {
		return response.ReturnCode(response.ErrNotFound)
	}

	return response.ReturnCode(response.CodeSuccess)
}
