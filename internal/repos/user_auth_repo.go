package repos

import (
	"fmt"
	"github.com/open-auth/global"
	"time"
)

type IUserAuthRepo interface {
	AddOTP(email string, opt int, expirationTime int64) error
}

type userAuthRepo struct{}

func NewUserAuthRepo() IUserAuthRepo {
	return &userAuthRepo{}
}

func (u *userAuthRepo) AddOTP(email string, opt int, expirationTime int64) error {
	key := fmt.Sprintf("user:%s:otp", email)
	return global.Rdb.SetEx(ctx, key, opt, time.Duration(expirationTime)).Err()
}
