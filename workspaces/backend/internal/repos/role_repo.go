package repos

import (
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/internal/db"
	"go.uber.org/zap"
)

type IRoleRepo interface {
	CreateNewRole(payload db.InsertNewRoleParams) error
	GetAllRoles() []db.Role
	GetById(id string) (*db.Role, error)
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

func (rr roleRepo) GetById(id string) (*db.Role, error) {
	role, err := rr.sqlC.GetRoleById(ctx, id)

	if err != nil {
		global.Logger.Error("GetById: ", zap.Error(err))
		return nil, err
	}

	return &role, nil
}
