package repos

import (
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/internal/db"
)

type ITokenRepo interface {
	CreateNewToken(payload db.CreateNewTokenParams) error
}

type tokenRepo struct {
	sqlC *db.Queries
}

func NewTokenRepo() ITokenRepo {
	return &tokenRepo{
		sqlC: db.New(global.Mdb),
	}
}

func (tr tokenRepo) CreateNewToken(payload db.CreateNewTokenParams) error {
	return tr.sqlC.CreateNewToken(ctx, payload)
}
