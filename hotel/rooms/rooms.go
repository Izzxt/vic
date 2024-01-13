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
	model   core.IRoomModel
	info    core.IRoomInfo
	tileMap core.IRoomTileMap
	habbos  []core.IHabbo
}

// GetHabbos implements core.IRoom.
func (r *Room) GetHabbos() []core.IHabbo {
	return r.habbos
}

// TileMap implements core.IRoom.
func (r *Room) TileMap() core.IRoomTileMap {
	return r.tileMap
}

// SetModel implements core.IRoom.
func (r *Room) SetModel(model core.IRoomModel) {
	r.model = model
}

// SuccessEnterRoom implements core.IRoom.
func (r *Room) SuccessEnterRoom(habbo core.IHabbo) {
	// habbo room unit
	counter++
	roomUnit := habbo_unit.NewHabboRoomUnit(counter, habbo, r, r.tileMap.GetDoorTile(), r.TileMap().GetDoorDirection())
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
func (r *Room) EnterRoom(habbo core.IHabbo) {
	if habbo == nil {
		habbo.Room().LeaveRoom(habbo, true)
	}

	if !isLoaded {
		r.LoadRoom()
	}

	r.PrepareRoom(habbo)
}

// LeaveRoom implements core.IRoom.
func (r *Room) LeaveRoom(habbo core.IHabbo, hotelview bool) {

	habbo.SetRoom(nil)

	// Remove habbo from room.

}

func (r *Room) GetHabbo(id int32) core.IHabbo {
	for _, habbo := range r.habbos {
		if habbo.HabboInfo().ID == id {
			return habbo
		}
	}
	return nil
}

func (r *Room) RemoveHabbo(habbo core.IHabbo) {
	for i, h := range r.habbos {
		if h.HabboInfo().ID == habbo.HabboInfo().ID {
			r.habbos = append(r.habbos[:i], r.habbos[i+1:]...)
			break
		}
	}
}

// PrepareRoom implements core.IRoom.
func (r *Room) PrepareRoom(habbo core.IHabbo) {
	habbo.SetRoom(r)
	r.habbos = append(r.habbos, habbo)

	habbo.Client().Send(&room.OpenComposer{})
	habbo.Client().Send(&room.ModelComposer{Room: r})

	// room paint composer
}

// UnloadRoom implements core.IRoom.
func (*Room) UnloadRoom(core.IHabbo) {
	panic("unimplemented")
}

func (r *Room) LoadRoom() {
	r.tileMap = tiles.NewRoomTileMap(r, r.model)

	isLoaded = true
}

// Model implements core.IRoom.
func (r *Room) Model() core.IRoomModel {
	return r.model
}

// Info implements core.IRoom.
func (r *Room) Info() core.IRoomInfo {
	return r.info
}

func NewRoom(ctx context.Context, roomInfo core.IRoomInfo) core.IRoom {
	room := Room{}
	room.info = roomInfo
	room.habbos = make([]core.IHabbo, 0)
	return &room
}
