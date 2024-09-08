package services

import (
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/internal/db"
	"github.com/go-open-auth/internal/dto"
	"github.com/go-open-auth/internal/repos"
	"github.com/go-open-auth/pkg/response"
	"github.com/go-open-auth/pkg/utils"
	"go.uber.org/zap"
)

type IRoleService interface {
	CreateNewRole(payload dto.RoleRequestDTO) *response.ServerCode
	GetAllRoles() []dto.RoleResponseDTO
	GetRole(id string) (*dto.RoleResponseDTO, *response.ServerCode)
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
	payloadRequest, errCode := utils.DtoToModel[db.InsertNewRoleParams](payload)

	if errCode != nil {
		return errCode
	}

	err := rs.roleRepo.CreateNewRole(*payloadRequest)

	if err != nil {
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
