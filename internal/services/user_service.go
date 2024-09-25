package services

import (
	"github.com/open-auth/internal/dto"
	"github.com/open-auth/internal/repos"
	"github.com/open-auth/pkg/response"
	"github.com/open-auth/pkg/utils"
)

type IUserService interface {
	GetMe(email string) (*dto.UserResponseDTO, *response.ServerCode)
	GetUsers(payload dto.SearchDTO) []dto.UserResponseDTO
}

type userService struct {
	userRepo repos.IUserRepo
}

func NewUserService(userRepo repos.IUserRepo) IUserService {
	return &userService{
		userRepo,
	}
}

func (us *userService) GetMe(email string) (*dto.UserResponseDTO, *response.ServerCode) {
	user, err := us.userRepo.GetUserByEmail(email)

	if err != nil {
		return nil, response.ReturnCode(response.ErrNotFound)
	}

	return utils.ModelToDto[dto.UserResponseDTO](*user), nil
}

func (us *userService) GetUsers(payload dto.SearchDTO) []dto.UserResponseDTO {
	return utils.ModelToDtos[dto.UserResponseDTO](
		us.userRepo.GetUsers(payload.Search, payload.By, payload.Limit(), payload.Skip()),
	)
}
