package habbo

import (
	"context"

	"github.com/Izzxt/vic/core"
	users "github.com/Izzxt/vic/database/users/querier"
)

type habbo struct {
	ctx       context.Context
	habboInfo users.User
	client    core.IHabboClient
	room      core.IRoom
	roomUnit  core.IHabboRoomUnit
}

// Client implements core.IHabbo.
func (h *habbo) Client() core.IHabboClient {
	return h.client
}

// Room implements core.IHabbo.
func (h *habbo) Room() core.IRoom {
	return h.room
}

func (h *habbo) SetRoom(room core.IRoom) {
	h.room = room
}

func (h *habbo) HabboInfo() users.User {
	return h.habboInfo
}

// RoomUnit implements core.IHabbo.
func (h *habbo) RoomUnit() core.IHabboRoomUnit {
	return h.roomUnit
}

// SetRoomUnit implements core.IHabbo.
func (h *habbo) SetRoomUnit(roomUnit core.IHabboRoomUnit) {
	h.roomUnit = roomUnit
}

func NewHabbo(ctx context.Context, users users.User, client core.IHabboClient) core.IHabbo {
	return &habbo{ctx: ctx, habboInfo: NewHabboInfo(ctx, users).User, client: client}
}
