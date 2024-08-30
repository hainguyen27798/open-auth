package services

import (
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/internal/dto"
	"github.com/go-open-auth/internal/repos"
	"github.com/go-open-auth/pkg/response"
	"github.com/go-open-auth/pkg/utils"
	"strconv"
	"time"
)

type IUserService interface {
	Register(user dto.UserRegistrationRequestDTO) int
	GetUsers() []string
}

type userService struct {
	userRepo     repos.IUserRepo
	userAuthRepo repos.IUserAuthRepo
}

func NewUserService(userRepo repos.IUserRepo, userAuthRepo repos.IUserAuthRepo) IUserService {
	return &userService{
		userRepo,
		userAuthRepo,
	}
}

func (us userService) GetUsers() []string {
	return us.userRepo.GetUsers()
}

func (us userService) Register(user dto.UserRegistrationRequestDTO) int {
	// hash email
	hashEmail := utils.GetHash(user.Email)

	// check email already exists
	if us.userRepo.CheckUserByEmail(user.Email) {
		return response.ErrCodeUserHasExists
	}

	// new OTP
	otp := utils.GenerateOTP()

	if err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute)); err != nil {
		return response.ErrInvalidOTP
	}

	// create user
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		global.Logger.Error(err.Error())
		return response.ErrCreateFailed
	}

	user.Password = hash
	if err := us.userRepo.CreateNewUser(user, strconv.Itoa(otp)); err != nil {
		global.Logger.Error(err.Error())
		return response.ErrCreateFailed
	}

	if err := utils.SendToEmail(
		"otp-email",
		"hainguyen27798@gmail.com",
		[]string{user.Email},
		map[string]interface{}{
			"otp": strconv.Itoa(otp),
		},
	); err != nil {
		return response.ErrSendEmailFailed
	}

	return response.ErrCodeSuccess
}
