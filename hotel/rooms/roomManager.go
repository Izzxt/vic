package rooms

import (
	"context"
	"fmt"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/database"
	room_info "github.com/Izzxt/vic/database/rooms/room_info/querier"
)

type RoomManager struct {
	ctx  context.Context
	room core.Room
}

var (
	rooms = make(map[int32]core.Room)
)

// Room implements core.IRoomManager.
func (r *RoomManager) GetRoom(id int32) core.Room {
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

func (r *RoomManager) GetActiveRooms() []room_info.GetActiveRoomsRow {
	db := database.GetInstance().RoomInfo()
	activeRooms, err := db.GetActiveRooms(r.ctx)
	if err != nil {
		fmt.Printf("failed to get active rooms: %v", err)
	}
	return activeRooms
}

func (r *RoomManager) GetRoomsByOwnerId(ownerId int32) []room_info.GetRoomsByOwnerIdRow {
	db := database.GetInstance().RoomInfo()
	roomInfos, err := db.GetRoomsByOwnerId(r.ctx, ownerId)
	if err != nil {
		fmt.Printf("failed to get rooms by owner id: %v", err)
	}
	return roomInfos
}

func NewRoomManager(ctx context.Context) core.RoomManager {
	roomManager := RoomManager{}
	roomManager.ctx = ctx
	return &roomManager
}
