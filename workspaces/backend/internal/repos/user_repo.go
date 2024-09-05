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

func (ur userRepo) GetUsers() []string {
	return []string{"hai", "harry"}
}

func (ur userRepo) CheckUserByEmail(email string) bool {
	_, err := ur.sqlC.GetUserByEmail(ctx, email)

	return !errors.Is(err, sql.ErrNoRows)
}

func (ur userRepo) CreateNewUser(userDto dto.UserRegistrationRequestDTO, code string) error {
	return ur.sqlC.CreateNewUser(ctx, db.CreateNewUserParams{
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

func (ur userRepo) GetUserById(email string) (*db.User, error) {
	user, err := ur.sqlC.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
