package rooms

import (
	"context"
	"sync"

	"github.com/Izzxt/vic/core"
	habbo_unit "github.com/Izzxt/vic/hotel/habbo/room"
	"github.com/Izzxt/vic/hotel/rooms/tiles"
	"github.com/Izzxt/vic/packets/outgoing/room"
	room_units "github.com/Izzxt/vic/packets/outgoing/room/units"
)

type Room struct {
	model    core.RoomModel
	info     core.RoomInfo
	tileMap  core.RoomTileMap
	isLoaded bool
	counter  int32
	habbos   []core.Habbo

	mu sync.RWMutex
}

// GetHabbos implements core.IRoom.
func (r *Room) GetHabbos() []core.Habbo {
	return r.habbos
}

// TileMap implements core.IRoom.
func (r *Room) TileMap() core.RoomTileMap {
	return r.tileMap
}

// SetModel implements core.IRoom.
func (r *Room) SetModel(model core.RoomModel) {
	r.mu.Lock()
	r.model = model
	r.mu.Unlock()
}

// SuccessEnterRoom implements core.IRoom.
func (r *Room) SuccessEnterRoom(habbo core.Habbo) {
	// habbo room unit
	roomUnit := habbo_unit.NewHabboRoomUnit(r.counter, habbo, r, r.tileMap.GetDoorTile(), r.TileMap().GetDoorDirection())
	r.counter++
	roomUnit.SetPreviousTile(r.tileMap.GetDoorTile())
	habbo.SetRoomUnit(roomUnit)
	habbo.Room().Info().UpdateOnlineCount(int32(len(r.habbos)))

	// send room user owner info
	// send room thickness
	// send room info
	habbo.Client().Send(&room.RoomDataComposer{Room: r, Enter: true, Forward: false})

	// send room unit
	habbo.Client().SendToRoom(r, &room_units.RoomUnitComposer{Habbos: r.habbos})
	// send room unit status
	habbo.Client().SendToRoom(r, room_units.NewRoomUnitStatusWithRoomsComposer(habbo.RoomUnit()))

}

// EnterRoom implements core.IRoom.
func (r *Room) EnterRoom(habbo core.Habbo) {
	if habbo.Room() != nil {
		habbo.Room().LeaveRoom(habbo, false)
	}

	if !r.isLoaded {
		r.LoadRoom()
	}

	r.PrepareRoom(habbo)
}

// LeaveRoom implements core.IRoom.
func (r *Room) LeaveRoom(habbo core.Habbo, hotelview bool) {
	habbo.Client().SendToRoom(r, &room_units.RoomUnitRemoveComposer{RoomUnit: habbo.RoomUnit()})

	habbo.SetRoom(nil)
	habbo.SetRoomUnit(nil)

	// Remove habbo from room.
	r.RemoveHabbo(habbo)

	// update users count
	r.info.UpdateOnlineCount(int32(len(r.habbos)))

	for _, t := range r.tileMap.GetDoorTile().HabboOnTiles() {
		t.CurrentTile().RemoveHabboOnTile(habbo)
	}

	if hotelview {
		// TODO: send to hotel view
	}

	if len(r.habbos) == 0 {
		r.UnloadRoom()
	}
}

func (r *Room) GetHabbo(id int32) core.Habbo {
	for _, habbo := range r.habbos {
		if habbo.HabboInfo().ID == id {
			return habbo
		}
	}
	return nil
}

func (r *Room) RemoveHabbo(habbo core.Habbo) {
	for i, h := range r.habbos {
		if h.HabboInfo().ID == habbo.HabboInfo().ID {
			r.habbos = append(r.habbos[:i], r.habbos[i+1:]...)
			break
		}
	}
}

// PrepareRoom implements core.IRoom.
func (r *Room) PrepareRoom(habbo core.Habbo) {
	habbo.SetRoom(r)
	r.habbos = append(r.habbos, habbo)

	habbo.Client().Send(&room.OpenComposer{})
	habbo.Client().Send(&room.ModelComposer{Room: r})

	// room paint composer
}

// UnloadRoom implements core.IRoom.
func (r *Room) UnloadRoom() {
	if !r.isLoaded {
		return
	}

	for _, habbo := range r.habbos {
		r.LeaveRoom(habbo, true)
	}

	r.tileMap = nil
	r.counter = 0
	r.isLoaded = false
}

func (r *Room) LoadRoom() {
	r.tileMap = tiles.NewRoomTileMap(r, r.model)

	r.isLoaded = true
}

// Model implements core.IRoom.
func (r *Room) Model() core.RoomModel {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.model
}

// Info implements core.IRoom.
func (r *Room) Info() core.RoomInfo {
	return r.info
}

func NewRoom(ctx context.Context, roomInfo core.RoomInfo, model core.RoomModel) core.Room {
	room := Room{}
	room.info = roomInfo
	room.isLoaded = false
	room.counter = 0
	room.model = model
	room.mu = sync.RWMutex{}
	room.habbos = make([]core.Habbo, 0)
	return &room
}
