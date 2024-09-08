package repos

import (
	"github.com/open-auth/global"
	"github.com/open-auth/internal/db"
	"go.uber.org/zap"
)

type IPermissionRepo interface {
	CreateNewPermission(payload db.InsertNewPermissionParams) error
	GetAllPermission() []db.Permission
	UpdatePermission(permission db.UpdatePermissionParams) (bool, error)
	DeletePermission(id string) bool
}

type permissionRepo struct {
	sqlC *db.Queries
}

func NewPermissionRepo() IPermissionRepo {
	return &permissionRepo{
		sqlC: db.New(global.Mdb),
	}
}

func (pr *permissionRepo) CreateNewPermission(payload db.InsertNewPermissionParams) error {
	return pr.sqlC.InsertNewPermission(ctx, payload)
}

func (pr *permissionRepo) GetAllPermission() []db.Permission {
	permission, err := pr.sqlC.GetAllPermissions(ctx)
	if err != nil {
		global.Logger.Error("GetAllPermission: ", zap.Error(err))
		return []db.Permission{}
	}
	return permission
}

func (pr *permissionRepo) UpdatePermission(permission db.UpdatePermissionParams) (bool, error) {
	affectRows, err := pr.sqlC.UpdatePermission(ctx, permission)

	if err != nil {
		return false, err
	}

	return affectRows > 0, nil
}

func (pr *permissionRepo) DeletePermission(id string) bool {
	count, err := pr.sqlC.DeletePermission(ctx, id)
	if err != nil {
		global.Logger.Error("DeletePermission: ", zap.Error(err))
		return false
	}
	return count > 0
}
