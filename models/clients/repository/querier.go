// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"context"
)

type Querier interface {
	CreateClient(ctx context.Context, arg CreateClientParams) (Client, error)
	DeleteUser(ctx context.Context, email string) (User, error)
	GetClientByEmail(ctx context.Context, email string) (GetClientByEmailRow, error)
}

var _ Querier = (*Queries)(nil)