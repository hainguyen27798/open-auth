package repos

import (
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/internal/db"
	"go.uber.org/zap"
)

type IRoleRepo interface {
	CreateNewRole(payload db.InsertNewRoleParams) error
	GetAllRoles() []db.Role
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

func (rr roleRepo) GetAllRoles() []db.Role {
	roles, err := rr.sqlC.GetAllRoles(ctx)

	if err != nil {
		global.Logger.Error("GetAllRoles: ", zap.Error(err))
		return []db.Role{}
	}

	return roles
}
