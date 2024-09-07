package repos

import (
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/internal/db"
)

type IPermissionRepo interface {
	CreateNewPermission(payload db.InsertNewPermissionParams) error
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
