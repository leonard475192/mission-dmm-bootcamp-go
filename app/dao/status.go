package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Status
	status struct {
		db *sqlx.DB
	}
)

// Create status repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// Create status :
func (r *status) Create(ctx context.Context, status object.Status) (*object.Status, error) {
	query, err := r.db.NamedExecContext(ctx, `INSERT INTO status (account_id, content) VALUES (:account_id,:content)`, status)
	if err != nil {
		return nil, fmt.Errorf("NamedExecContext:%w", err)
	}
	id, err := query.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("LastInsertId:%w", err)
	}

	// StructScan でリレーションやりたかった
	// TODO QueryRowxContext:missing destination name username in *object.Status
	// entity := new(object.Status)
	// log.Print(entity)
	// err = r.db.QueryRowxContext(ctx, comfirm, id).StructScan(entity)
	// if err != nil {
	// 	return nil, fmt.Errorf("QueryRowxContext:%w", err)
	// }
	var statuses []object.Status
	const comfirm = `
		SELECT
			status.id AS id,
			status.content AS content,
			status.create_at AS create_at,
			account.id AS "account.id",
			account.username AS "account.username",
			account.create_at AS "account.create_at"
		FROM status
		JOIN account ON status.account_id = account.id
		WHERE status.id = ?
	`
	err = r.db.SelectContext(ctx, &statuses, comfirm, id)
	if err != nil {
		return nil, fmt.Errorf("SelectContext:%w", err)
	}

	return &statuses[0], nil
}

func (r *status) Get(ctx context.Context, id string) (*object.Status, error) {
	var statuses []object.Status
	const comfirm = `
		SELECT
			status.id AS id,
			status.content AS content,
			status.create_at AS create_at,
			account.id AS "account.id",
			account.username AS "account.username",
			account.create_at AS "account.create_at"
		FROM status
		JOIN account ON status.account_id = account.id
		WHERE status.id = ?
	`
	err := r.db.SelectContext(ctx, &statuses, comfirm, id)
	if err != nil {
		return nil, fmt.Errorf("SelectContext:%w", err)
	}

	return &statuses[0], nil
}
