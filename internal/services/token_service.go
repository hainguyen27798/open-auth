package services

import (
	"fmt"
	"github.com/go-open-auth/internal/db"
	"github.com/go-open-auth/internal/repos"
	"github.com/go-open-auth/pkg/response"
	"github.com/go-open-auth/pkg/utils"
)

type ITokenService interface {
	GenerateNewToken(user db.User) (*utils.Token, error)
	ReNewToken(token string) (*utils.Token, *response.ServerCode)
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

	token, err := utils.GenerateJWT(user.ID, map[string]interface{}{
		"name":    user.Name,
		"email":   user.Email,
		"session": session,
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

func (ts *tokenService) ReNewToken(token string) (*utils.Token, *response.ServerCode) {
	if ts.tokenRepo.CheckOldRefreshTokenExists(token) {
		return nil, response.ReturnCode(response.ErrStolenToken)
	}

	claims, errCode := utils.VerifyJWT(token)
	if errCode != nil {
		return nil, errCode
	}

	newToken, err := utils.GenerateJWT(claims.UserID, claims.Data)
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
