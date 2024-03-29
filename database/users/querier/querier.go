// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package users

import (
	"context"
)

type Querier interface {
	GetUserByAuthTicket(ctx context.Context, authTicket string) (User, error)
	ListUsers(ctx context.Context) ([]User, error)
}

var _ Querier = (*Queries)(nil)
