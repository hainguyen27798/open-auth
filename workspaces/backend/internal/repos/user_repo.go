package repos

import (
	"database/sql"
	"errors"
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/internal/db"
)

type IUserRepo interface {
	CheckUserByEmail(email string) bool
	GetUsers() []string
}

type userRepo struct{}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}

func (ur userRepo) GetUsers() []string {
	return []string{"hai", "harry"}
}

func (ur userRepo) CheckUserByEmail(email string) bool {
	q := db.New(global.Mdb)
	_, err := q.GetUserByEmail(ctx, email)

	return !errors.Is(err, sql.ErrNoRows)
}
