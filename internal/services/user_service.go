package services

import (
	"github.com/open-auth/internal/db"
	"github.com/open-auth/internal/repos"
	"github.com/open-auth/pkg/response"
)

type IUserService interface {
	GetMe(email string) (*db.User, *response.ServerCode)
}

type userService struct {
	userRepo repos.IUserRepo
}

func NewUserService(userRepo repos.IUserRepo) IUserService {
	return &userService{
		userRepo,
	}
}

func (us *userService) GetMe(email string) (*db.User, *response.ServerCode) {
	user, err := us.userRepo.GetUserById(email)

	if err != nil {
		return nil, response.ReturnCode(response.ErrNotFound)
	}

	return user, nil
}
