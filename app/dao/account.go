package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// Create User :
func (r *account) Create(ctx context.Context, account object.Account) (*object.Account, error) {
	query, err := r.db.NamedExecContext(ctx, `INSERT INTO account (username, password_hash) VALUES (:username,:password_hash)`, account)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	id, err := query.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	entity := new(object.Account)
	err = r.db.QueryRowxContext(ctx, "SELECT * FROM account WHERE id = ?", id).StructScan(entity)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}
