package repos

import (
	"database/sql"
	"errors"
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/internal/db"
	"github.com/go-open-auth/internal/dto"
	"github.com/google/uuid"
)

type IUserRepo interface {
	CheckUserByEmail(email string) bool
	CreateNewUser(userDto dto.UserRegistrationRequestDTO, code string) error
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

func (ur userRepo) CreateNewUser(userDto dto.UserRegistrationRequestDTO, code string) error {
	q := db.New(global.Mdb)

	return q.CreateNewUser(ctx, db.CreateNewUserParams{
		ID:               uuid.New().String(),
		Name:             userDto.Name,
		Email:            userDto.Email,
		Password:         sql.NullString{String: userDto.Password, Valid: true},
		Status:           db.UsersStatusRequest,
		SocialProvider:   db.NullUsersSocialProvider{},
		Image:            sql.NullString{},
		VerificationCode: sql.NullString{String: code, Valid: true},
	})
}
