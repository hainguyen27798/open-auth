package repos

import (
	"github.com/jmoiron/sqlx"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/models"
	"github.com/open-auth/internal/query"
	"github.com/open-auth/pkg/utils"
	"go.uber.org/zap"
)

type IUserRepo interface {
	CheckUserByEmail(email string) bool
	CreateNewUser(userDto models.InsertNewUserParams) error
	CreateSuperUser(adminDto models.InsertSuperUserParams) error
	SearchUsers(search string, by string, limit int, skip int) ([]models.User, int64)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByEmailAndScope(email string, scope models.UsersScope) (*models.User, error)
}

type userRepo struct {
	sqlX *sqlx.DB
}

func NewUserRepo() IUserRepo {
	return &userRepo{
		sqlX: global.Mdb,
	}
}

func (ur *userRepo) SearchUsers(search string, by string, limit int, skip int) ([]models.User, int64) {
	var users []models.User
	var total int64
	queryString := query.SearchUserBy[by]
	queryCount := query.CountSearchUserBy[by]
	search = "%" + search + "%"

	if queryString == "" {
		queryString = query.SearchUserBy["name"]
		queryCount = query.CountSearchUserBy["name"]
	}

	if err := ur.sqlX.Select(&users, queryString, search, limit, skip); err != nil {
		global.Logger.Error("CountUser: ", zap.Error(err))
		return []models.User{}, 0
	}

	if err := ur.sqlX.Get(&total, queryCount, search); err != nil {
		global.Logger.Error("CountUser: ", zap.Error(err))
		return []models.User{}, 0
	}

	return users, total
}

func (ur *userRepo) CheckUserByEmail(email string) bool {
	var exists bool
	if err := ur.sqlX.Get(&exists, query.CheckUserByEmail, email); err != nil {
		return false
	}
	return exists
}

func (ur *userRepo) CreateNewUser(payload models.InsertNewUserParams) error {
	session, err := utils.NewTransaction(ur.sqlX)
	if err != nil {
		return err
	}

	if _, err := session.NamedExecCommit(query.InsertBasicUser, payload); err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) CreateSuperUser(payload models.InsertSuperUserParams) error {
	session, err := utils.NewTransaction(ur.sqlX)
	if err != nil {
		return err
	}

	if _, err := session.NamedExecCommit(query.InsertSuperuser, payload); err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := ur.sqlX.Get(&user, query.GetUserByEmail, email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepo) GetUserByEmailAndScope(email string, scope models.UsersScope) (*models.User, error) {
	var user models.User

	if err := ur.sqlX.Get(&user, query.GetUserByEmailAndScope, email, scope); err != nil {
		return nil, err
	}

	return &user, nil
}
