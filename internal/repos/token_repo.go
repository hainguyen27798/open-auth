package repos

import (
	"database/sql"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/db"
)

type ITokenRepo interface {
	CreateNewToken(payload db.CreateNewTokenParams) error
	UpdateRefreshToken(session string, newRefreshToken string) error
	CheckOldRefreshTokenExists(oldRefreshToken string) bool
	RemoveToken(token string) bool
}

type tokenRepo struct {
	sqlC *db.Queries
}

func NewTokenRepo() ITokenRepo {
	return &tokenRepo{
		sqlC: db.New(global.Mdb),
	}
}

func (tr *tokenRepo) CreateNewToken(payload db.CreateNewTokenParams) error {
	return tr.sqlC.CreateNewToken(ctx, payload)
}

func (tr *tokenRepo) UpdateRefreshToken(session string, newRefreshToken string) error {
	tx, err := global.Mdb.Begin()
	if err != nil {
		return err
	}

	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
		}
	}(tx)

	q := db.New(tx)

	token, err := q.GetTokenBySession(ctx, session)
	if err != nil {
		return err
	}

	if err := q.UpdateRefreshToken(ctx, db.UpdateRefreshTokenParams{
		RefreshToken: newRefreshToken,
		ID:           token.ID,
	}); err != nil {
		return err
	}

	if err := q.CacheOldRefreshToken(ctx, db.CacheOldRefreshTokenParams{
		TokenID:      token.ID,
		RefreshToken: token.RefreshToken,
	}); err != nil {
		return err
	}

	return tx.Commit()
}

func (tr *tokenRepo) CheckOldRefreshTokenExists(oldRefreshToken string) bool {
	count, _ := tr.sqlC.CheckOldRefreshTokenExists(ctx, oldRefreshToken)
	return count > 0
}

func (tr *tokenRepo) RemoveToken(token string) bool {
	affectRow, err := tr.sqlC.RemoveToken(ctx, token)
	if err != nil {
		return false
	}
	return affectRow > 0
}
