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
	Register(user dto.UserRegistrationRequestDTO) *response.ServerCode
	Login(user dto.UserLoginRequestDTO) (*dto.UserLoginResponseDTO, *response.ServerCode)
	RefreshToken(token string) (*dto.TokenResponseDTO, *response.ServerCode)
	Logout(token string) *response.ServerCode
}

type authService struct {
	userRepo     repos.IUserRepo
	userAuthRepo repos.IUserAuthRepo
	tokenService ITokenService
}

func NewAuthService(userRepo repos.IUserRepo, userAuthRepo repos.IUserAuthRepo, tokenService ITokenService) IAuthService {
	return &authService{
		userRepo,
		userAuthRepo,
		tokenService,
	}
}

func (as *authService) Register(user dto.UserRegistrationRequestDTO) *response.ServerCode {
	// hash email
	hashEmail := utils.GetHash(user.Email)

	// check email already exists
	if as.userRepo.CheckUserByEmail(user.Email) {
		return response.ReturnCode(response.ErrCodeUserHasExists)
	}

	// new OTP
	otp := utils.GenerateOTP()

	if err := as.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute)); err != nil {
		return response.ReturnCode(response.ErrInvalidOTP)
	}

	// create user
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		global.Logger.Error(err.Error())
		return response.ReturnCode(response.ErrCreateFailed)
	}

	user.Password = hash
	if err := as.userRepo.CreateNewUser(user, strconv.Itoa(otp)); err != nil {
		global.Logger.Error(err.Error())
		return response.ReturnCode(response.ErrCreateFailed)
	}

	if err := utils.SendToEmail(
		"otp-email",
		"hainguyen27798@gmail.com",
		[]string{user.Email},
		map[string]interface{}{
			"otp": strconv.Itoa(otp),
		},
	); err != nil {
		return response.ReturnCode(response.ErrSendEmailFailed)
	}

	return response.ReturnCode(response.CodeSuccess)
}

func (as *authService) Login(user dto.UserLoginRequestDTO) (*dto.UserLoginResponseDTO, *response.ServerCode) {
	userExisting, err := as.userRepo.GetUserById(user.Email)
	if err != nil {
		return nil, response.ReturnCode(response.ErrCodeUserNotExists)
	}

	if utils.VerifyPassword(user.Password, userExisting.Password.String) {

		token, err := as.tokenService.GenerateNewToken(*userExisting)
		if err != nil {
			return nil, response.ReturnCode(response.ErrJWTInternalError)
		}

		return &dto.UserLoginResponseDTO{
			ID:    userExisting.ID,
			Name:  userExisting.Name,
			Email: userExisting.Email,
			TokenResponseDTO: dto.TokenResponseDTO{
				AccessToken:  token.AccessToken,
				RefreshToken: token.RefreshToken,
			},
		}, nil
	}

	return nil, response.ReturnCode(response.ErrCodeLoginFailed)
}

func (as *authService) RefreshToken(token string) (*dto.TokenResponseDTO, *response.ServerCode) {
	if newToken, errCode := as.tokenService.ReNewToken(token); errCode != nil {
		return nil, errCode
	} else {
		return &dto.TokenResponseDTO{
			AccessToken:  newToken.AccessToken,
			RefreshToken: newToken.RefreshToken,
		}, nil
	}
}

func (as *authService) Logout(token string) *response.ServerCode {
	return as.tokenService.RemoveToken(token)
}
