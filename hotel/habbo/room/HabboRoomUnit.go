package habbo

import (
	"fmt"
	"time"

	"github.com/Izzxt/vic/core"
	room_units "github.com/Izzxt/vic/packets/outgoing/room/units"
)

type habboRoomUnit struct {
	id    int32
	habbo core.IHabbo
	room  core.IRoom

	currentTile  core.IRoomTile
	previousTile core.IRoomTile
	headRotation core.RoomTileDirection
	bodyRotation core.RoomTileDirection
	goal         core.IRoomTile
	start        core.IRoomTile
	statuses     map[core.HabboRoomUnitStatus]string
}

// Statuses implements core.IHabboRoomUnit.
func (h *habboRoomUnit) Statuses() map[core.HabboRoomUnitStatus]string {
	return h.statuses
}

// PreviousTile implements core.IHabboRoomUnit.
func (h *habboRoomUnit) PreviousTile() core.IRoomTile {
	return h.previousTile
}

// SetPreviousTile implements core.IHabboRoomUnit.
func (h *habboRoomUnit) SetPreviousTile(tile core.IRoomTile) {
	h.previousTile = tile
}

// WalkTo implements core.IHabboRoomUnit.
func (h *habboRoomUnit) WalkTo(tile core.IRoomTile) {
	if h.room == nil {
		return
	}

	h.SetPreviousTile(tile)

	dest := h.Room().TileMap().FindPath(h.CurrentTile(), h.previousTile)
	if len(dest) == 0 {
		return
	}

	h.statuses = make(map[core.HabboRoomUnitStatus]string)

	next := dest[0]
	count := len(dest)
	for range dest {
		count--

		time.Sleep(500 * time.Millisecond)
		h.SetPreviousTile(dest[count])
		h.SetCurrentTile(dest[count])
		h.statuses[core.HabboRoomUnitStatus(core.HabboRoomUnitStatusMove)] = fmt.Sprintf("%d,%d,%d", next.GetX(), next.GetY(), next.GetHeight())
		go h.habbo.Client().SendToRoom(h.room, &room_units.RoomUnitStatusComposer{Habbos: []core.IHabbo{h.habbo}})
	}
}

func (h *habboRoomUnit) ID() int32 {
	return h.id
}

func (h *habboRoomUnit) Habbo() core.IHabbo {
	return h.habbo
}

func (h *habboRoomUnit) Room() core.IRoom {
	return h.room
}

func (h *habboRoomUnit) CurrentTile() core.IRoomTile {
	return h.currentTile
}

func (h *habboRoomUnit) HeadRotation() core.RoomTileDirection {
	return h.headRotation
}

func (h *habboRoomUnit) BodyRotation() core.RoomTileDirection {
	return h.bodyRotation
}

func (h *habboRoomUnit) SetCurrentTile(tile core.IRoomTile) {
	h.currentTile = tile
}

func (h *habboRoomUnit) SetHeadRotation(rotation core.RoomTileDirection) {
	h.headRotation = rotation
}

func (h *habboRoomUnit) SetBodyRotation(rotation core.RoomTileDirection) {
	h.bodyRotation = rotation
}

func (h *habboRoomUnit) SetRoom(room core.IRoom) {
	h.room = room
}

func (h *habboRoomUnit) SetHabbo(habbo core.IHabbo) {
	h.habbo = habbo
}

func (h *habboRoomUnit) Dispose() {
	h.room = nil
	h.habbo = nil
	h.currentTile = nil
}

func NewHabboRoomUnit(id int32, habbo core.IHabbo, room core.IRoom, currentTile core.IRoomTile, bodyRotation core.RoomTileDirection) core.IHabboRoomUnit {
	habboRoomUnit := new(habboRoomUnit)
	habboRoomUnit.id = id
	habboRoomUnit.habbo = habbo
	habboRoomUnit.room = room
	habboRoomUnit.currentTile = currentTile
	habboRoomUnit.bodyRotation = bodyRotation
	currentTile.AddHabboRoomUnit(habboRoomUnit)
	return habboRoomUnit
}
