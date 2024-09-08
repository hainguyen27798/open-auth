package services

import (
	"github.com/go-open-auth/internal/db"
	"github.com/go-open-auth/internal/dto"
	"github.com/go-open-auth/internal/repos"
	"github.com/go-open-auth/pkg/response"
	"github.com/go-open-auth/pkg/utils"
)

type IRoleService interface {
	CreateNewRole(payload dto.RoleRequestDTO) *int
	GetAllRole() []dto.RoleResponseDTO
}

type roleService struct {
	roleRepo repos.IRoleRepo
}

func NewRoleService(roleRepo repos.IRoleRepo) IRoleService {
	return &roleService{
		roleRepo,
	}
}

func (rs *roleService) CreateNewRole(payload dto.RoleRequestDTO) *int {
	payloadRequest, errCode := utils.DtoToModel[db.InsertNewRoleParams](payload)

	if errCode != nil {
		return errCode
	}

	err := rs.roleRepo.CreateNewRole(*payloadRequest)

	if err != nil {
		return &[]int{response.ErrCreateFailed}[0]
	}

	return &[]int{response.CreatedSuccess}[0]
}

func (rs *roleService) GetAllRole() []dto.RoleResponseDTO {
	return utils.ModelToDtos[dto.RoleResponseDTO](rs.roleRepo.GetAllRoles())
}
