package services

import (
	"github.com/go-open-auth/internal/db"
	"github.com/go-open-auth/internal/repos"
	"github.com/go-open-auth/pkg/response"
)

type IUserService interface {
	GetMe(email string) (*db.User, *int)
}

type userService struct {
	userRepo repos.IUserRepo
}

func NewUserService(userRepo repos.IUserRepo) IUserService {
	return userService{
		userRepo,
	}
}

func (us userService) GetMe(email string) (*db.User, *int) {
	user, err := us.userRepo.GetUserById(email)

	if err != nil {
		return nil, &[]int{response.ErrNotFound}[0]
	}

	return user, nil
}
