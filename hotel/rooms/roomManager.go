package rooms

import (
	"context"
	"fmt"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/database"
	room_info "github.com/Izzxt/vic/database/rooms/room_info/querier"
)

type RoomManager struct {
	ctx   context.Context
	room  core.Room
	model core.RoomModel
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
	model := NewRoomModels(r.ctx).Load(roomInfo.GetModelId())
	room := NewRoom(r.ctx, roomInfo, model)
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

func (r *RoomManager) CreateRoom(ownerId int32, name string, description string, modelId int32, categoryId int32, maxVisitors int32, tradeType int32) core.Room {
	db := database.GetInstance().RoomInfo()
	id, err := db.CreateRoom(r.ctx, room_info.CreateRoomParams{
		OwnerID:        ownerId,
		Name:           name,
		Description:    description,
		ModelID:        modelId,
		FlatCategoryID: categoryId,
		MaxUsers:       maxVisitors,
		TradeMode:      tradeType,
	})
	if err != nil {
		fmt.Printf("failed to create room: %v", err)
	}

	room := r.GetRoom(int32(id))

	return room
}

func (r *RoomManager) Shutdown() {
	for _, room := range rooms {
		room.UnloadRoom()
	}
}

func (r *RoomManager) Model() core.RoomModel {
	return r.model
}

func NewRoomManager(ctx context.Context) core.RoomManager {
	roomManager := RoomManager{}
	roomManager.ctx = ctx
	return &roomManager
}
