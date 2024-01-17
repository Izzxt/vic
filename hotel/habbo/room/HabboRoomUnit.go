package habbo

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/hotel/rooms/tiles"
	"github.com/Izzxt/vic/list"
	room_units "github.com/Izzxt/vic/packets/outgoing/room/units"
)

type habboRoomUnit struct {
	id    int32
	habbo core.Habbo
	room  core.Room

	currentTile  core.IRoomTile
	previousTile core.IRoomTile
	headRotation core.RoomTileDirection
	bodyRotation core.RoomTileDirection
	goalTile     core.IRoomTile
	goalPath     list.List[core.IRoomTile]
	statuses     map[core.HabboRoomUnitStatus]string
	statusMutex  sync.Mutex
	t            *time.Timer
	ticker       *time.Ticker

	goalTileMutex    sync.Mutex
	goalPathMutex    sync.Mutex
	currentTileMutex sync.Mutex
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
func (h *habboRoomUnit) WalkTo(ctx context.Context, tile core.IRoomTile, client core.HabboClient) {
	if h.room == nil {
		return
	}

	h.goalTileMutex.Lock()
	h.goalTile = tile
	h.goalTileMutex.Unlock()

	if h.ticker == nil {
		ctx, cancel := context.WithCancel(ctx)

		delay := 500 * time.Millisecond
		h.ticker = PereodicallyDo(ctx, delay, func(ctx context.Context, ticker *time.Ticker, _ time.Time, wg *sync.WaitGroup) {
			defer wg.Done()
			h.goalPath = h.Room().TileMap().FindPath(h.currentTile, h.goalTile).Pop()

			next := h.goalPath.Last()

			if h.goalTile == nil {
				h.stopWalking(cancel)
				return
			}

			if h.goalPath.IsEmpty() {
				h.stopWalking(cancel)
				return
			}

			if next == h.currentTile {
				h.goalPath = h.goalPath.Pop()
			}

			direction := tiles.GetRoomTileDirection(h.CurrentTile(), next)
			h.SetBodyRotation(direction)
			h.SetHeadRotation(direction)

			h.statusMutex.Lock()
			h.statuses[core.HabboRoomUnitStatus(core.HabboRoomUnitStatusMove)] = fmt.Sprintf("%d,%d,%.1f", next.GetX(), next.GetY(), next.GetHeight())
			h.statusMutex.Unlock()

			h.SetPreviousTile(h.currentTile)
			h.SetCurrentTile(next)

			client.SendToRoom(h.room, room_units.NewRoomUnitStatusWithHabbosComposer(client.GetHabbo().RoomUnit().Room().GetHabbos()))
		})
	}
}

func (h *habboRoomUnit) stopWalking(cancel context.CancelFunc) {
	h.SetPreviousTile(h.currentTile)

	delete(h.statuses, core.HabboRoomUnitStatus(core.HabboRoomUnitStatusMove))

	h.habbo.Client().SendToRoom(h.room, room_units.NewRoomUnitStatusWithHabbosComposer(h.Room().GetHabbos()))

	h.ticker.Stop()
	cancel()
	h.ticker = nil

	h.goalTile = nil
	h.goalPath = nil
}

func (h *habboRoomUnit) ID() int32 {
	return h.id
}

func (h *habboRoomUnit) Habbo() core.Habbo {
	return h.habbo
}

func (h *habboRoomUnit) Room() core.Room {
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

func (h *habboRoomUnit) SetRoom(room core.Room) {
	h.room = room
}

func (h *habboRoomUnit) SetHabbo(habbo core.Habbo) {
	h.habbo = habbo
}

func (h *habboRoomUnit) Dispose() {
	h.room = nil
	h.habbo = nil
	h.currentTile = nil
}

func NewHabboRoomUnit(id int32, habbo core.Habbo, room core.Room, currentTile core.IRoomTile, bodyRotation core.RoomTileDirection) core.HabboRoomUnit {
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

func PereodicallyDo(ctx context.Context, delay time.Duration, f func(ctx context.Context, ticker *time.Ticker, time time.Time, wg *sync.WaitGroup)) *time.Ticker {
	wg := sync.WaitGroup{}
	ticker := time.NewTicker(delay)
	go func() {
		for {
			wg.Add(1)
			select {
			case <-ctx.Done():
				return
			case now := <-ticker.C:
				f(ctx, ticker, now, &wg)
			}
		}
	}()
	wg.Wait()

	return ticker
}
