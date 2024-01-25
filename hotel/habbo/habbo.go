package habbo

import (
	"context"
	"sync"
	"unicode/utf8"

	"github.com/Izzxt/vic/core"
	users "github.com/Izzxt/vic/database/users/querier"
	users_stats "github.com/Izzxt/vic/database/users/stats/querier"
	room_chat "github.com/Izzxt/vic/packets/outgoing/room/units/chats"
)

type habbo struct {
	ctx       context.Context
	habboInfo users.User
	client    core.HabboClient
	room      core.Room
	roomUnit  core.HabboRoomUnit
	stats     core.HabboStats

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

func (h *habbo) HabboStats() core.HabboStats {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.stats
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

// Whisper implements core.IHabbo.
func (h *habbo) Whisper(to core.Habbo, message string, gesture int32) {
	h.client.Send(&room_chat.RoomUnitChatWhisperComposer{
		RoomUnit: to.RoomUnit(), Message: message, Gesture: gesture,
		Bubble: h.HabboStats().GetBubbleChat(), MessageLength: int32(utf8.RuneCountInString(message))})
}

// Shout implements core.IHabbo.
func (h *habbo) Shout(message string, gesture int32) {
	h.client.SendToRoom(h.Room(), &room_chat.RoomUnitChatShoutComposer{
		RoomUnit: h.RoomUnit(), Message: message, Gesture: gesture,
		Bubble: h.HabboStats().GetBubbleChat(), MessageLength: int32(utf8.RuneCountInString(message))})
}

// Talk implements core.IHabbo.
func (h *habbo) Talk(message string, gesture int32) {
	h.client.SendToRoom(h.Room(), &room_chat.RoomUnitChatComposer{
		RoomUnit: h.RoomUnit(), Message: message, Gesture: gesture,
		Bubble: h.HabboStats().GetBubbleChat(), MessageLength: int32(utf8.RuneCountInString(message))})
}

func NewHabbo(ctx context.Context, users users.User, client core.HabboClient) core.Habbo {
	return &habbo{
		ctx: ctx, habboInfo: NewHabboInfo(ctx, users).User, client: client,
		stats: NewHabboStats(ctx, users_stats.UsersStat{}).Load(users.ID),
	}
}
