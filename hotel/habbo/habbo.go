package habbo

import (
	"context"
	"sync"

	"github.com/Izzxt/vic/core"
	users "github.com/Izzxt/vic/database/users/querier"
)

type habbo struct {
	ctx       context.Context
	habboInfo users.User
	client    core.HabboClient
	room      core.Room
	roomUnit  core.HabboRoomUnit

	mu sync.RWMutex
}

// Client implements core.IHabbo.
func (h *habbo) Client() core.HabboClient {
	return h.client
}

// Room implements core.IHabbo.
func (h *habbo) Room() core.Room {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.room
}

func (h *habbo) SetRoom(room core.Room) {
	h.mu.Lock()
	h.room = room
	h.mu.Unlock()
}

func (h *habbo) HabboInfo() users.User {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.habboInfo
}

// RoomUnit implements core.IHabbo.
func (h *habbo) RoomUnit() core.HabboRoomUnit {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.roomUnit
}

// SetRoomUnit implements core.IHabbo.
func (h *habbo) SetRoomUnit(roomUnit core.HabboRoomUnit) {
	h.mu.Lock()
	h.roomUnit = roomUnit
	h.mu.Unlock()
}

func NewHabbo(ctx context.Context, users users.User, client core.HabboClient) core.Habbo {
	return &habbo{ctx: ctx, habboInfo: NewHabboInfo(ctx, users).User, client: client}
}
