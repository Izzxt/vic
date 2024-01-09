package tiles

import "github.com/Izzxt/vic/core"

type RoomTile struct {
	x, y, height int32
	state        core.RoomTileState
	habboOnTiles []core.IHabboRoomUnit
}

// GetHabboRoomUnit implements core.IRoomTile.
func (r *RoomTile) HabboOnTiles() []core.IHabboRoomUnit {
	return r.habboOnTiles
}

func (r *RoomTile) AddHabboRoomUnit(h core.IHabboRoomUnit) {
	r.habboOnTiles = append(r.habboOnTiles, h)
}

func (r *RoomTile) RemoveHabboRoomUnit(h core.IHabboRoomUnit) {
	for i, habbo := range r.habboOnTiles {
		if habbo == h {
			r.habboOnTiles = append(r.habboOnTiles[:i], r.habboOnTiles[i+1:]...)
		}
	}
}

// GetHeight implements core.IRoomTile.
func (r *RoomTile) GetHeight() int32 {
	return r.height
}

// GetState implements core.IRoomTile.
func (r *RoomTile) GetState() core.RoomTileState {
	return r.state
}

// GetX implements core.IRoomTile.
func (r *RoomTile) GetX() int32 {
	return r.x
}

// GetY implements core.IRoomTile.
func (r *RoomTile) GetY() int32 {
	return r.y
}

func NewRoomTile(x int32, y int32, height int32, state core.RoomTileState) core.IRoomTile {
	return &RoomTile{x: x, y: y, height: height, state: state}
}
