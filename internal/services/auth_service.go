package services

import (
	"context"
	"encoding/json"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/dto"
	"github.com/open-auth/internal/models"
	"github.com/open-auth/internal/repos"
	"github.com/open-auth/pkg/response"
	"github.com/open-auth/pkg/utils"
	"github.com/segmentio/kafka-go"
	"strconv"
	"strings"
	"time"
)

type IAuthService interface {
	Register(user dto.UserRegistrationRequestDTO) *response.ServerCode
	Login(user dto.UserLoginRequestDTO, scope global.Scope) (*dto.UserLoginResponseDTO, *response.ServerCode)
	RefreshToken(scope global.Scope, token string) (*dto.TokenResponseDTO, *response.ServerCode)
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
	payload, errCode := utils.DtoToModel[models.InsertNewUserParams](user)
	payload.VerificationCode = strconv.Itoa(otp)
	if errCode != nil {
		return errCode
	}

	if err := as.userRepo.CreateNewUser(*payload); err != nil {
		global.Logger.Error(err.Error())
		return response.ReturnCode(response.ErrCreateFailed)
	}

	body := make(map[string]interface{})
	body["templateName"] = "otp-email"
	body["subject"] = "OTP Verification"
	body["to"] = []string{user.Email}
	body["data"] = map[string]string{"otp": strconv.Itoa(otp)}
	message, _ := json.Marshal(body)
	go func(msg []byte) {
		err := global.SMTPProducer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte("opt-auth"),
			Value: msg,
			Time:  time.Now(),
		})
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		global.Logger.Info("Send OTP to " + user.Email)
	}(message)

	return response.ReturnCode(response.CodeSuccess)
}

func (as *authService) Login(user dto.UserLoginRequestDTO, scope global.Scope) (*dto.UserLoginResponseDTO, *response.ServerCode) {
	userExisting, err := as.userRepo.GetUserByEmailAndScope(user.Email, models.UsersScope(strings.ToLower(string(scope))))
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

func (as *authService) RefreshToken(scope global.Scope, token string) (*dto.TokenResponseDTO, *response.ServerCode) {
	if newToken, errCode := as.tokenService.ReNewToken(scope, token); errCode != nil {
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
