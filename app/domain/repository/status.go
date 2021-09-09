package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Fetch account which has specified username
	// TODO: Add Other APIs
	Create(ctx context.Context, status object.Status) (*object.Status, error)
}
