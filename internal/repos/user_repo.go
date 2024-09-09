package repos

import (
	"database/sql"
	"errors"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/db"
)

type IUserRepo interface {
	CheckUserByEmail(email string) bool
	CreateNewUser(userDto db.InsertNewUserParams) error
	CreateSuperUser(adminDto db.InsertSuperUserParams) error
	GetUsers() []string
	GetUserById(email string) (*db.User, error)
}

type userRepo struct {
	sqlC *db.Queries
}

func NewUserRepo() IUserRepo {
	return &userRepo{
		sqlC: db.New(global.Mdb),
	}
}

func (ur *userRepo) GetUsers() []string {
	return []string{"hai", "harry"}
}

func (ur *userRepo) CheckUserByEmail(email string) bool {
	_, err := ur.sqlC.GetUserByEmail(ctx, email)

	return !errors.Is(err, sql.ErrNoRows)
}

func (ur *userRepo) CreateNewUser(payload db.InsertNewUserParams) error {
	return ur.sqlC.InsertNewUser(ctx, payload)
}

func (ur *userRepo) CreateSuperUser(payload db.InsertSuperUserParams) error {
	return ur.sqlC.InsertSuperUser(ctx, payload)
}

func (ur *userRepo) GetUserById(email string) (*db.User, error) {
	user, err := ur.sqlC.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
