package repos

import (
	"github.com/jmoiron/sqlx"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/models"
	"github.com/open-auth/internal/query"
	"github.com/open-auth/pkg/utils"
)

type ITokenRepo interface {
	CreateNewToken(payload models.InsertNewTokenParams) error
	UpdateRefreshToken(session string, newRefreshToken string) error
	CheckOldRefreshTokenExists(oldRefreshToken string) bool
	RemoveToken(token string) bool
}

type tokenRepo struct {
	sqlX *sqlx.DB
}

func NewTokenRepo() ITokenRepo {
	return &tokenRepo{
		sqlX: global.MdbX,
	}
}

func (tr *tokenRepo) CreateNewToken(payload models.InsertNewTokenParams) error {
	session, err := utils.NewTransaction(tr.sqlX)
	if err != nil {
		return err
	}

	if _, err := session.NamedExecCommit(query.InsertNewToken, payload); err != nil {
		return err
	}

	return nil
}

func (tr *tokenRepo) UpdateRefreshToken(session string, newRefreshToken string) error {
	var token models.Token
	if err := tr.sqlX.Get(&token, query.GetTokenBySession, session); err != nil {
		return err
	}

	tran, err := utils.NewTransaction(tr.sqlX)
	if err != nil {
		return err
	}

	tran.Exec(query.UpdateRefreshToken, newRefreshToken, token.ID)
	tran.Exec(query.CacheOldRefreshToken, token.ID, token.RefreshToken)

	tran.Commit()

	return nil
}

func (tr *tokenRepo) CheckOldRefreshTokenExists(oldRefreshToken string) bool {
	var exists bool
	if err := tr.sqlX.Get(&exists, query.CheckOldRefreshTokenExists, oldRefreshToken); err != nil {
		return false
	}
	return exists
}

func (tr *tokenRepo) RemoveToken(token string) bool {
	session, err := utils.NewTransaction(tr.sqlX)
	if err != nil {
		return false
	}

	count, err := session.ExecCommit(query.RemoveToken, token)
	if err != nil {
		return false
	}
	return count > 0
}
