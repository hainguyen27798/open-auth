package services

import (
	"github.com/go-open-auth/internal/db"
	"github.com/go-open-auth/internal/repos"
	"github.com/go-open-auth/pkg/utils"
	"github.com/google/uuid"
)

type ITokenService interface {
	GenerateNewToken(user db.User) (*utils.Token, error)
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
		ID:           uuid.New().String(),
		UserID:       user.ID,
		Session:      session,
		RefreshToken: token.RefreshToken,
	}); err != nil {
		return nil, err
	}

	return token, nil
}
