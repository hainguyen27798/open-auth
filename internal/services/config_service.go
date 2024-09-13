package services

import (
	"database/sql"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/db"
	"github.com/open-auth/internal/repos"
	"github.com/open-auth/pkg/utils"
)

type IConfigService interface {
	InitAdmin(email string, password string)
}

type configService struct {
	userRepo repos.IUserRepo
}

func NewConfigService() IConfigService {
	return &configService{
		userRepo: repos.NewUserRepo(),
	}
}

func (cs configService) InitAdmin(email string, password string) {
	if ok := cs.userRepo.CheckUserByEmail(email); ok {
		global.Logger.Info("Admin is already initialized")
		return
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		global.Logger.Error(err.Error())
		panic(err)
		return
	}

	if err := cs.userRepo.CreateSuperUser(db.InsertSuperUserParams{
		Email: email,
		Password: sql.NullString{
			String: hash,
			Valid:  true,
		},
	}); err != nil {
		global.Logger.Error(err.Error())
		panic(err)
		return
	}

	global.Logger.Info("Admin initialized")
}
