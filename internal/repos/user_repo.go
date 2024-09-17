package repos

import (
	"github.com/jmoiron/sqlx"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/models"
	"github.com/open-auth/internal/query"
	"github.com/open-auth/pkg/utils"
)

type IUserRepo interface {
	CheckUserByEmail(email string) bool
	CreateNewUser(userDto models.InsertNewUserParams) error
	CreateSuperUser(adminDto models.InsertSuperUserParams) error
	GetUsers() []string
	GetUserByEmail(email string) (*models.User, error)
	GetUserByEmailAndScope(email string, scope models.UsersScope) (*models.User, error)
}

type userRepo struct {
	sqlX *sqlx.DB
}

func NewUserRepo() IUserRepo {
	return &userRepo{
		sqlX: global.MdbX,
	}
}

func (ur *userRepo) GetUsers() []string {
	return []string{"hai", "harry"}
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
