package habbo

import (
	"context"

	"github.com/Izzxt/vic/core"
	users "github.com/Izzxt/vic/database/users/querier"
)

type habbo struct {
	ctx       context.Context
	habboInfo users.User
	client    core.HabboClient
	room      core.Room
	roomUnit  core.HabboRoomUnit
}

// Client implements core.IHabbo.
func (h *habbo) Client() core.HabboClient {
	return h.client
}

// Room implements core.IHabbo.
func (h *habbo) Room() core.Room {
	return h.room
}

func (h *habbo) SetRoom(room core.Room) {
	h.room = room
}

func (h *habbo) HabboInfo() users.User {
	return h.habboInfo
}

// RoomUnit implements core.IHabbo.
func (h *habbo) RoomUnit() core.HabboRoomUnit {
	return h.roomUnit
}

// SetRoomUnit implements core.IHabbo.
func (h *habbo) SetRoomUnit(roomUnit core.HabboRoomUnit) {
	h.roomUnit = roomUnit
}

func NewHabbo(ctx context.Context, users users.User, client core.HabboClient) core.Habbo {
	return &habbo{ctx: ctx, habboInfo: NewHabboInfo(ctx, users).User, client: client}
}
