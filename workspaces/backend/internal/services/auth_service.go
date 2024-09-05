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

type IAuthService interface {
	Register(user dto.UserRegistrationRequestDTO) int
	Login(user dto.UserLoginRequestDTO) (*dto.UserLoginResponseDTO, *int)
}

type authService struct {
	userRepo     repos.IUserRepo
	userAuthRepo repos.IUserAuthRepo
}

func NewAuthService(userRepo repos.IUserRepo, userAuthRepo repos.IUserAuthRepo) IAuthService {
	return &authService{
		userRepo,
		userAuthRepo,
	}
}

func (as authService) Register(user dto.UserRegistrationRequestDTO) int {
	// hash email
	hashEmail := utils.GetHash(user.Email)

	// check email already exists
	if as.userRepo.CheckUserByEmail(user.Email) {
		return response.ErrCodeUserHasExists
	}

	// new OTP
	otp := utils.GenerateOTP()

	if err := as.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute)); err != nil {
		return response.ErrInvalidOTP
	}

	// create user
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		global.Logger.Error(err.Error())
		return response.ErrCreateFailed
	}

	user.Password = hash
	if err := as.userRepo.CreateNewUser(user, strconv.Itoa(otp)); err != nil {
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

	return response.CodeSuccess
}

func (as authService) Login(user dto.UserLoginRequestDTO) (*dto.UserLoginResponseDTO, *int) {
	userExisting, err := as.userRepo.GetUserById(user.Email)
	if err != nil {
		errCode := response.ErrCodeUserNotExists
		return nil, &errCode
	}

	if utils.VerifyPassword(user.Password, userExisting.Password.String) {
		return &dto.UserLoginResponseDTO{
			ID:           userExisting.ID,
			Name:         userExisting.Name,
			Email:        userExisting.Email,
			AccessToken:  "",
			RefreshToken: "",
		}, nil
	}

	errCode := response.ErrCodeLoginFailed
	return nil, &errCode
}
