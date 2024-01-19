// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package room_info

import (
	"context"
)

type Querier interface {
	GetActiveRooms(ctx context.Context) ([]GetActiveRoomsRow, error)
	GetRoomById(ctx context.Context, id int32) (Room, error)
	GetRoomsByOwnerId(ctx context.Context, id int32) ([]GetRoomsByOwnerIdRow, error)
	ListRooms(ctx context.Context) ([]Room, error)
}

var _ Querier = (*Queries)(nil)
