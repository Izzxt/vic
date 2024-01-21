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

	mu sync.RWMutex

	currentTile core.RoomTile
	ctMu        sync.RWMutex

	previousTile core.RoomTile
	ptMu         sync.RWMutex

	headRotation core.RoomTileDirection
	hrMu         sync.RWMutex

	bodyRotation core.RoomTileDirection
	brMu         sync.RWMutex

	goalTile core.RoomTile
	gtMu     sync.RWMutex

	goalPath list.List[core.RoomTile]
	gpMu     sync.RWMutex

	statuses map[core.HabboRoomUnitStatus]string
	statMu   sync.RWMutex

	ticker *time.Ticker
	tkMu   sync.RWMutex
}

// p	panic("unimplemented")tatuses implements core.IHabboRoomUnit.
func (h *habboRoomUnit) Statuses() map[core.HabboRoomUnitStatus]string {
	h.statMu.Lock()
	defer h.statMu.Unlock()

	return h.statuses
}

// PreviousTile implements core.IHabboRoomUnit.
func (h *habboRoomUnit) PreviousTile() core.RoomTile {
	h.ptMu.RLock()
	defer h.ptMu.RUnlock()

	return h.previousTile
}

// SetPreviousTile implements core.IHabboRoomUnit.
func (h *habboRoomUnit) SetPreviousTile(tile core.RoomTile) {
	h.ptMu.Lock()
	h.previousTile = tile
	h.ptMu.Unlock()
}

// WalkTo implements core.IHabboRoomUnit.
func (h *habboRoomUnit) WalkTo(ctx context.Context, tile core.RoomTile, client core.HabboClient) {
	if h.Room() == nil || h.Habbo() == nil {
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	h.goalTile = tile

	if h.ticker == nil {
		ctx, cancel := context.WithCancel(ctx)

		delay := 500 * time.Millisecond
		h.ticker = PereodicallyDo(ctx, delay, func(ctx context.Context, ticker *time.Ticker, _ time.Time, wg *sync.WaitGroup) {
			if h.Room() == nil || h.Habbo() == nil || h.Room().TileMap() == nil {
				h.stopWalking(cancel)
				return
			}

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

			h.statuses[core.HabboRoomUnitStatus(core.HabboRoomUnitStatusMove)] = fmt.Sprintf("%d,%d,%.1f", next.GetX(), next.GetY(), next.GetHeight())

			h.SetPreviousTile(h.currentTile)
			h.SetCurrentTile(next)

			client.SendToRoom(h.room, room_units.NewRoomUnitStatusWithHabbosComposer([]core.Habbo{h.Habbo()}))
		})
	}
}

func (h *habboRoomUnit) stopWalking(cancel context.CancelFunc) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.SetPreviousTile(h.currentTile)

	delete(h.statuses, core.HabboRoomUnitStatus(core.HabboRoomUnitStatusMove))
	h.habbo.Client().SendToRoom(h.room, room_units.NewRoomUnitStatusWithHabbosComposer([]core.Habbo{h.habbo}))

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

func (h *habboRoomUnit) CurrentTile() core.RoomTile {
	h.ctMu.RLock()
	defer h.ctMu.RUnlock()

	return h.currentTile
}

func (h *habboRoomUnit) HeadRotation() core.RoomTileDirection {
	h.hrMu.RLock()
	defer h.hrMu.RUnlock()

	return h.headRotation
}

func (h *habboRoomUnit) BodyRotation() core.RoomTileDirection {
	h.brMu.RLock()
	defer h.brMu.RUnlock()

	return h.bodyRotation
}

func (h *habboRoomUnit) SetCurrentTile(tile core.RoomTile) {
	h.ctMu.Lock()
	h.currentTile = tile
	h.ctMu.Unlock()
}

func (h *habboRoomUnit) SetHeadRotation(rotation core.RoomTileDirection) {
	h.hrMu.Lock()
	h.headRotation = rotation
	h.hrMu.Unlock()
}

func (h *habboRoomUnit) SetBodyRotation(rotation core.RoomTileDirection) {
	h.brMu.Lock()
	h.bodyRotation = rotation
	h.brMu.Unlock()
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

func NewHabboRoomUnit(id int32, habbo core.Habbo, room core.Room, currentTile core.RoomTile, bodyRotation core.RoomTileDirection) core.HabboRoomUnit {
	habboRoomUnit := new(habboRoomUnit)
	habboRoomUnit.id = id
	habboRoomUnit.habbo = habbo
	habboRoomUnit.room = room
	habboRoomUnit.currentTile = currentTile
	habboRoomUnit.bodyRotation = bodyRotation
	habboRoomUnit.statMu = sync.RWMutex{}
	habboRoomUnit.ptMu = sync.RWMutex{}
	habboRoomUnit.hrMu = sync.RWMutex{}
	habboRoomUnit.brMu = sync.RWMutex{}
	habboRoomUnit.gtMu = sync.RWMutex{}
	habboRoomUnit.gpMu = sync.RWMutex{}
	habboRoomUnit.ctMu = sync.RWMutex{}
	habboRoomUnit.tkMu = sync.RWMutex{}
	habboRoomUnit.mu = sync.RWMutex{}
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
