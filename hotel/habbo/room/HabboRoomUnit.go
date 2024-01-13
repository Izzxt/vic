package habbo

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/hotel/rooms/tiles"
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
	statusMutex  sync.Mutex
}

// p	panic("unimplemented")tatuses implements core.IHabboRoomUnit.
func (h *habboRoomUnit) Statuses() map[core.HabboRoomUnitStatus]string {
	h.statusMutex.Lock()
	defer h.statusMutex.Unlock()
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
func (h *habboRoomUnit) WalkTo(ctx context.Context, tile core.IRoomTile) {
	if h.room == nil {
		return
	}
	ctx, cancel := context.WithCancel(ctx)

	h.SetPreviousTile(tile)

	algo := h.Room().TileMap().FindPath(h.CurrentTile(), h.previousTile)

	delay := time.NewTicker(250 * time.Millisecond)
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			<-delay.C
			if algo.Len() == 0 {
				delete(h.statuses, core.HabboRoomUnitStatus(core.HabboRoomUnitStatusMove))
				cancel()
			}

			select {
			case <-ticker.C:

				next := algo.Last()

				if next == h.currentTile {
					algo.Pop()
					continue
				}

				fmt.Println(next.GetX(), next.GetY(), next.GetHeight())

				direction := tiles.GetRoomTileDirection(h.CurrentTile(), next)
				h.SetBodyRotation(direction)
				h.SetHeadRotation(direction)

				h.statusMutex.Lock()
				h.statuses[core.HabboRoomUnitStatus(core.HabboRoomUnitStatusMove)] = fmt.Sprintf("%d,%d,%.1f", next.GetX(), next.GetY(), next.GetHeight())
				h.statusMutex.Unlock()

				h.SetPreviousTile(algo.Last())

				h.SetCurrentTile(next)

				algo.Pop()

				go h.habbo.Client().SendToRoom(h.room, room_units.NewRoomUnitStatusWithHabbosComposer([]core.IHabbo{h.habbo}))

			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}

	}()
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
	habboRoomUnit.statusMutex = sync.Mutex{}
	habboRoomUnit.statuses = make(map[core.HabboRoomUnitStatus]string)
	currentTile.AddHabboRoomUnit(habboRoomUnit)
	return habboRoomUnit
}
