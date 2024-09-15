package utils

import (
	"github.com/jmoiron/sqlx"
	"github.com/open-auth/global"
	"go.uber.org/zap"
)

type Transaction struct {
	tx *sqlx.Tx
}

func NewTransaction(db *sqlx.DB) (*Transaction, error) {
	tx, err := db.Beginx()
	if err != nil {
		global.Logger.Error("Failed to start transaction", zap.Error(err))
		return nil, err
	}
	return &Transaction{
		tx,
	}, nil
}

func (t *Transaction) Commit() {
	err := t.tx.Commit()
	if err != nil {
		global.Logger.Error("Failed to commit transaction", zap.Error(err))
		t.Rollback()
		return
	}
}

func (t *Transaction) Rollback() {
	err := t.tx.Rollback()
	global.Logger.Debug("Rollback transaction")
	if err != nil {
		global.Logger.Error("Failed to rollback transaction", zap.Error(err))
		return
	}
}

func (t *Transaction) NamedExec(query string, args interface{}) (bool, error) {
	rs, err := t.tx.NamedExec(query, args)
	if err != nil {
		t.Rollback()
		return false, err
	}

	rowsAffected, err := rs.RowsAffected()
	if err != nil {
		t.Rollback()
		return false, err
	}

	t.Commit()
	return rowsAffected > 0, nil
}

func (t *Transaction) Exec(query string, args ...interface{}) (int64, error) {
	rs := t.tx.MustExec(query, args...)
	t.Commit()
	return rs.RowsAffected()
}
