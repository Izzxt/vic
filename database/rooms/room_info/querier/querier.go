// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package room_info

import (
	"context"
)

type Querier interface {
	GetRoomById(ctx context.Context, id int32) (Room, error)
	ListRooms(ctx context.Context) ([]Room, error)
}

var _ Querier = (*Queries)(nil)