package repos

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/models"
	"github.com/open-auth/internal/query"
	"github.com/open-auth/pkg/utils"
	"go.uber.org/zap"
)

type IRoleRepo interface {
	CreateNewRole(payload models.InsertNewRoleParams) error
	GetAllRoles(search string, skip int, limit int) ([]models.Role, int64)
	GetById(id string) (*models.Role, error)
	Delete(id string) bool
	Update(payload models.UpdateRoleParams) (bool, error)
	InsertRolePermission(roleId string, permissionId string) error
	GetRolePermissions(roleId string) []models.Permission
}

type roleRepo struct {
	sqlX *sqlx.DB
}

func NewRoleRepo() IRoleRepo {
	return &roleRepo{
		sqlX: global.Mdb,
	}
}

func (rr *roleRepo) CreateNewRole(payload models.InsertNewRoleParams) error {
	session, err := utils.NewTransaction(rr.sqlX)
	if err != nil {
		return err
	}

	if _, err := session.NamedExecCommit(query.InsertNewRole, payload); err != nil {
		return err
	}

	return nil
}

func (rr *roleRepo) GetAllRoles(search string, skip int, limit int) ([]models.Role, int64) {
	var roles []models.Role
	var total int64

	if err := rr.sqlX.Select(&roles, query.GetAllRoles, "%"+search+"%", limit, skip); err != nil {
		global.Logger.Error("GetAllRoles: ", zap.Error(err))
		return []models.Role{}, 0
	}

	if err := rr.sqlX.Get(&total, query.GetAllRolesTotal, "%"+search+"%"); err != nil {
		global.Logger.Error("CountPermission: ", zap.Error(err))
		return []models.Role{}, 0
	}

	return roles, total
}

func (rr *roleRepo) GetById(id string) (*models.Role, error) {
	var role models.Role

	err := rr.sqlX.Get(&role, query.GetRoleById, id)

	if err != nil {
		global.Logger.Error("GetById: ", zap.Error(err))
		return nil, err
	}

	return &role, nil
}

func (rr *roleRepo) Delete(id string) bool {
	session, err := utils.NewTransaction(rr.sqlX)
	if err != nil {
		return false
	}

	count, err := session.ExecCommit(query.DeleteRole, id)
	if err != nil {
		return false
	}
	return count > 0
}

func (rr *roleRepo) Update(payload models.UpdateRoleParams) (bool, error) {
	role, err := rr.GetById(*payload.ID)

	if err != nil {
		return false, err
	}

	querySet := utils.PartialUpdate(payload)

	if role == nil || querySet == "" {
		return false, nil
	}

	queryString := fmt.Sprintf(query.UpdateRole, "SET "+querySet)

	session, err := utils.NewTransaction(rr.sqlX)
	if err != nil {
		return false, err
	}

	_, err = session.NamedExecCommit(queryString, payload)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (rr *roleRepo) InsertRolePermission(roleId string, permissionId string) error {
	session, err := utils.NewTransaction(rr.sqlX)
	if err != nil {
		return err
	}
	session.Exec(query.InsertRolePermission, roleId, permissionId)
	session.Commit()
	return nil
}

func (rr *roleRepo) GetRolePermissions(roleId string) []models.Permission {
	var permissions []models.Permission
	err := rr.sqlX.Select(&permissions, query.SelectPermissionByRoleId, roleId)
	if err != nil {
		global.Logger.Error("GetRolePermissions: ", zap.Error(err))
		return []models.Permission{}
	}
	return permissions
}
