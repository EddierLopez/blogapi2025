package dto

import (
	"context"
	"database/sql"
	"fmt"
)

type DbTransaction struct {
	*Queries
	db *sql.DB
}

func NewDbTransaction(db *sql.DB) *DbTransaction {
	return &DbTransaction{
		db:      db,
		Queries: New(db),
	}
}

func (dbTransaction *DbTransaction) ExcTransaction(ctx context.Context, fn func(*Queries) error) error {
	tx, err := dbTransaction.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbError := tx.Rollback(); rbError != nil {
			return fmt.Errorf("tx Error:%v, Rollback error:%v", err, rbError)
		}
		return err
	}
	return tx.Commit()
}
