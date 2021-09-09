package dao

import (
	"context"
	"fmt"
	"log"
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

	// TODO QueryRowxContext:missing destination name username in *object.Status
	const comfirm = `
		SELECT *
		FROM status
		JOIN account ON status.account_id = account.id
		WHERE status.id = ?
	`

	entity := new(object.Status)
	log.Print(entity)
	err = r.db.QueryRowxContext(ctx, comfirm, id).StructScan(entity)
	if err != nil {
		return nil, fmt.Errorf("QueryRowxContext:%w", err)
	}
	// var statuses []object.Status
	// err = r.db.SelectContext(ctx, &statuses, "SELECT * FROM status JOIN account ON status.account_id = account.id WHERE status.id = ?", id)
	// if err != nil {
	// 	return nil, fmt.Errorf("SelectContext:%w", err)
	// }

	return entity, nil
}
