package repos

import (
	"github.com/open-auth/global"
	"github.com/open-auth/internal/db"
	"go.uber.org/zap"
)

type IRoleRepo interface {
	CreateNewRole(payload db.InsertNewRoleParams) error
	GetAllRoles() []db.Role
	GetById(id string) (*db.Role, error)
	Delete(id string) (bool, error)
	Update(payload db.UpdateRoleParams) error
}

type roleRepo struct {
	sqlC *db.Queries
}

func NewRoleRepo() IRoleRepo {
	return &roleRepo{
		sqlC: db.New(global.Mdb),
	}
}

func (rr *roleRepo) CreateNewRole(payload db.InsertNewRoleParams) error {
	return rr.sqlC.InsertNewRole(ctx, payload)
}

func (rr *roleRepo) GetAllRoles() []db.Role {
	roles, err := rr.sqlC.GetAllRoles(ctx)

	if err != nil {
		global.Logger.Error("GetAllRoles: ", zap.Error(err))
		return []db.Role{}
	}

	return roles
}

func (rr *roleRepo) GetById(id string) (*db.Role, error) {
	role, err := rr.sqlC.GetRoleById(ctx, id)

	if err != nil {
		global.Logger.Error("GetById: ", zap.Error(err))
		return nil, err
	}

	return &role, nil
}

func (rr *roleRepo) Delete(id string) (bool, error) {
	affectRows, err := rr.sqlC.DeleteRole(ctx, id)

	if err != nil {
		global.Logger.Error("DeleteRole: ", zap.Error(err))
		return false, err
	}

	return affectRows > 0, nil
}

func (rr *roleRepo) Update(payload db.UpdateRoleParams) error {
	if err := rr.sqlC.UpdateRole(ctx, payload); err != nil {
		global.Logger.Error("UpdateRole: ", zap.Error(err))
		return err
	}
	return nil
}
