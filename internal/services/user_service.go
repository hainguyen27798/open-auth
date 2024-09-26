package services

import (
	"github.com/open-auth/internal/dto"
	"github.com/open-auth/internal/repos"
	"github.com/open-auth/pkg/response"
	"github.com/open-auth/pkg/utils"
)

type IUserService interface {
	GetMe(email string) (*dto.UserResponseDTO, *response.ServerCode)
	SearchUsers(payload dto.SearchDTO) dto.PaginationDto[dto.UserResponseDTO]
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

func (us *userService) SearchUsers(payload dto.SearchDTO) dto.PaginationDto[dto.UserResponseDTO] {
	users, total := us.userRepo.SearchUsers(payload.Search, payload.By, payload.Limit(), payload.Skip())
	return utils.ModelToPaginationDto[dto.UserResponseDTO](
		users,
		dto.PaginationMetaDataDto{
			Total:        total,
			PageSize:     payload.Limit(),
			PageSelected: payload.PageSelected(),
		},
	)
}
