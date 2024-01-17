package rooms

import (
	"context"

	"github.com/Izzxt/vic/core"
	habbo_unit "github.com/Izzxt/vic/hotel/habbo/room"
	"github.com/Izzxt/vic/hotel/rooms/tiles"
	"github.com/Izzxt/vic/packets/outgoing/room"
	room_units "github.com/Izzxt/vic/packets/outgoing/room/units"
)

var (
	isLoaded       = false
	counter  int32 = 0
)

type Room struct {
	model   core.RoomModel
	info    core.RoomInfo
	tileMap core.RoomTileMap
	habbos  []core.Habbo
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
	r.model = model
}

// SuccessEnterRoom implements core.IRoom.
func (r *Room) SuccessEnterRoom(habbo core.Habbo) {
	// habbo room unit
	roomUnit := habbo_unit.NewHabboRoomUnit(counter, habbo, r, r.tileMap.GetDoorTile(), r.TileMap().GetDoorDirection())
	counter++
	roomUnit.SetPreviousTile(r.tileMap.GetDoorTile())
	habbo.SetRoomUnit(roomUnit)
	// update room habbos size

	// send room user owner info
	// send room thickness
	// send room info

	// send room unit
	habbo.Client().SendToRoom(r, &room_units.RoomUnitComposer{Habbos: r.habbos})
	// send room unit status
	habbo.Client().SendToRoom(r, room_units.NewRoomUnitStatusWithRoomsComposer(habbo.RoomUnit()))

}

// EnterRoom implements core.IRoom.
func (r *Room) EnterRoom(habbo core.Habbo) {
	if habbo == nil {
		habbo.Room().LeaveRoom(habbo, true)
	}

	if !isLoaded {
		r.LoadRoom()
	}

	r.PrepareRoom(habbo)
}

// LeaveRoom implements core.IRoom.
func (r *Room) LeaveRoom(habbo core.Habbo, hotelview bool) {

	habbo.SetRoom(nil)

	// Remove habbo from room.

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
func (*Room) UnloadRoom(core.Habbo) {
	panic("unimplemented")
}

func (r *Room) LoadRoom() {
	r.tileMap = tiles.NewRoomTileMap(r, r.model)

	isLoaded = true
}

// Model implements core.IRoom.
func (r *Room) Model() core.RoomModel {
	return r.model
}

// Info implements core.IRoom.
func (r *Room) Info() core.RoomInfo {
	return r.info
}

func NewRoom(ctx context.Context, roomInfo core.RoomInfo) core.Room {
	room := Room{}
	room.info = roomInfo
	room.habbos = make([]core.Habbo, 0)
	return &room
}
