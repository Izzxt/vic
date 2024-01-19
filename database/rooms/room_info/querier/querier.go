// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package room_info

import (
	"context"
)

type Querier interface {
	CreateRoom(ctx context.Context, arg CreateRoomParams) (int64, error)
	GetActiveRooms(ctx context.Context) ([]GetActiveRoomsRow, error)
	GetRoomById(ctx context.Context, id int32) (GetRoomByIdRow, error)
	GetRoomsByOwnerId(ctx context.Context, id int32) ([]GetRoomsByOwnerIdRow, error)
	ListRooms(ctx context.Context) ([]Room, error)
	UpdateRoomUsers(ctx context.Context, arg UpdateRoomUsersParams) error
}

var _ Querier = (*Queries)(nil)
