package repos

import (
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/internal/db"
	"go.uber.org/zap"
)

type IPermissionRepo interface {
	CreateNewPermission(payload db.InsertNewPermissionParams) error
	GetAllPermission() []db.Permission
	UpdatePermission(permission db.UpdatePermissionParams) error
}

type permissionRepo struct {
	sqlC *db.Queries
}

func NewPermissionRepo() IPermissionRepo {
	return &permissionRepo{
		sqlC: db.New(global.Mdb),
	}
}

func (pr permissionRepo) CreateNewPermission(payload db.InsertNewPermissionParams) error {
	return pr.sqlC.InsertNewPermission(ctx, payload)
}

func (pr permissionRepo) GetAllPermission() []db.Permission {
	permission, err := pr.sqlC.GetAllPermissions(ctx)
	if err != nil {
		global.Logger.Error("GetAllPermission: ", zap.Error(err))
		return []db.Permission{}
	}
	return permission
}

func (pr permissionRepo) UpdatePermission(permission db.UpdatePermissionParams) error {
	return pr.sqlC.UpdatePermission(ctx, permission)
}
