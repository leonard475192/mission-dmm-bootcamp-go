package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	// TODO: Add Other APIs
	Create(ctx context.Context, user object.Account) (*object.Account, error)
}
