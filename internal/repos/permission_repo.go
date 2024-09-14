package repos

import (
	"github.com/jmoiron/sqlx"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/db"
	"github.com/open-auth/internal/models"
	"github.com/open-auth/internal/sql"
	"go.uber.org/zap"
)

type IPermissionRepo interface {
	CreateNewPermission(payload db.InsertNewPermissionParams) error
	GetAllPermission(search string, by string) []models.Permission
	UpdatePermission(permission db.UpdatePermissionParams) (bool, error)
	DeletePermission(id string) bool
}

type permissionRepo struct {
	sqlC *db.Queries
	sqlX *sqlx.DB
}

func NewPermissionRepo() IPermissionRepo {
	return &permissionRepo{
		sqlC: db.New(global.Mdb),
		sqlX: global.MdbX,
	}
}

func (pr *permissionRepo) CreateNewPermission(payload db.InsertNewPermissionParams) error {
	return pr.sqlC.InsertNewPermission(ctx, payload)
}

func (pr *permissionRepo) GetAllPermission(search string, by string) []models.Permission {
	var permission []models.Permission
	query := sql.GetAllPermissionsBy[by]
	search += "%"

	if search != "" {
		query = sql.GetAllPermissionsBy["service_name"]
	}

	if err := pr.sqlX.Select(&permission, query, search); err != nil {
		global.Logger.Error("GetAllPermission: ", zap.Error(err))
		return []models.Permission{}
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
