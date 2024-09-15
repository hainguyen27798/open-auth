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
	GetAllPermission(search string, by string, skip int, limit int) ([]models.Permission, int64)
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

func (pr *permissionRepo) GetAllPermission(search string, by string, skip int, limit int) ([]models.Permission, int64) {
	var permission []models.Permission
	var total int64
	query := sql.GetAllPermissionsBy[by]
	queryCount := sql.CountPermissionSearchBy[by]
	search = "%" + search + "%"

	if query == "" {
		query = sql.GetAllPermissionsBy["service_name"]
		queryCount = sql.CountPermissionSearchBy["service_name"]
	}

	if err := pr.sqlX.Select(&permission, query, search, limit, skip); err != nil {
		global.Logger.Error("GetAllPermission: ", zap.Error(err))
		return []models.Permission{}, 0
	}

	if err := pr.sqlX.Get(&total, queryCount, search); err != nil {
		global.Logger.Error("CountPermission: ", zap.Error(err))
		return []models.Permission{}, 0
	}

	return permission, total
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
