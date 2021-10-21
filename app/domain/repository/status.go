package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Fetch account which has specified username
	// TODO: Add Other APIs
	Create(ctx context.Context, status object.Status) (*object.Status, error)
	Get(ctx context.Context, id int) (*object.Status, error)
	Select(ctx context.Context, only_media bool, greater_than_id int, less_than_id int, limit int) ([]object.Status, error)
	Delete(ctx context.Context, id int) error
}
