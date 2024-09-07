package repos

import (
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/internal/db"
)

type IRoleRepo interface {
	CreateNewRole(payload db.InsertNewRoleParams) error
}

type roleRepo struct {
	sqlC *db.Queries
}

func NewRoleRepo() IRoleRepo {
	return &roleRepo{
		sqlC: db.New(global.Mdb),
	}
}

func (rr roleRepo) CreateNewRole(payload db.InsertNewRoleParams) error {
	return rr.sqlC.InsertNewRole(ctx, payload)
}
