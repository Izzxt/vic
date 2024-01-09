package rooms

import (
	"context"

	"github.com/Izzxt/vic/core"
)

type RoomManager struct {
	ctx  context.Context
	room core.IRoom
}

var (
	rooms = make(map[int32]core.IRoom)
)

// Room implements core.IRoomManager.
func (r *RoomManager) GetRoom(id int32) core.IRoom {
	if room, ok := rooms[id]; ok {
		return room
	}

	roomInfo := NewRoomInfo(r.ctx).Load(id)
	room := NewRoom(r.ctx, roomInfo)

	model := NewRoomModels(r.ctx).Load(roomInfo.GetModelId())
	room.SetModel(model)
	rooms[id] = room

	return room
}

func NewRoomManager(ctx context.Context) core.IRoomManager {
	roomManager := RoomManager{}
	roomManager.ctx = ctx
	return &roomManager
}
