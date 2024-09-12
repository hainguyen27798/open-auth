package services

import (
	"fmt"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/db"
	"github.com/open-auth/internal/repos"
	"github.com/open-auth/pkg/response"
	"github.com/open-auth/pkg/utils"
	"strings"
)

type ITokenService interface {
	GenerateNewToken(user db.User) (*utils.Token, error)
	ReNewToken(scope global.Scope, token string) (*utils.Token, *response.ServerCode)
	RemoveToken(token string) *response.ServerCode
}

type tokenService struct {
	tokenRepo repos.ITokenRepo
}

func NewTokenService(tokenRepo repos.ITokenRepo) ITokenService {
	return &tokenService{
		tokenRepo,
	}
}

func (ts *tokenService) GenerateNewToken(user db.User) (*utils.Token, error) {
	session := utils.CreateSession(32)
	scope := global.Scope(strings.ToUpper(string(user.Scope)))

	token, err := utils.GenerateJWT(scope, user.ID, map[string]interface{}{
		"name":    user.Name,
		"email":   user.Email,
		"session": session,
		"scope":   scope,
	})
	if err != nil {
		return nil, err
	}

	if err := ts.tokenRepo.CreateNewToken(db.CreateNewTokenParams{
		UserID:       user.ID,
		Session:      session,
		RefreshToken: token.RefreshToken,
	}); err != nil {
		return nil, err
	}

	return token, nil
}

func (ts *tokenService) ReNewToken(scope global.Scope, token string) (*utils.Token, *response.ServerCode) {
	if ts.tokenRepo.CheckOldRefreshTokenExists(token) {
		return nil, response.ReturnCode(response.ErrStolenToken)
	}

	claims, errCode := utils.VerifyJWT(scope, token)
	if errCode != nil {
		return nil, errCode
	}

	newToken, err := utils.GenerateJWT(scope, claims.UserID, claims.Data)
	if err != nil {
		return nil, response.ReturnCode(response.ErrJWTInternalError)
	}

	session := fmt.Sprintf("%v", claims.Data["session"])

	if err := ts.tokenRepo.UpdateRefreshToken(session, newToken.RefreshToken); err != nil {
		return nil, response.ReturnCode(response.ErrJWTInternalError)
	}

	return newToken, nil
}

func (ts *tokenService) RemoveToken(token string) *response.ServerCode {
	if ok := ts.tokenRepo.RemoveToken(token); ok {
		return response.ReturnCode(response.LogoutSuccess)
	}
	return response.ReturnCode(response.ErrInvalidToken)
}
