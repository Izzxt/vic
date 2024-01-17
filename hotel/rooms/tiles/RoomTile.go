package tiles

import "github.com/Izzxt/vic/core"

type RoomTile struct {
	x, y         int32
	height       float32
	state        core.RoomTileState
	habboOnTiles []core.HabboRoomUnit
}

// GetHabboRoomUnit implements core.IRoomTile.
func (r *RoomTile) HabboOnTiles() []core.HabboRoomUnit {
	return r.habboOnTiles
}

func (r *RoomTile) AddHabboRoomUnit(h core.HabboRoomUnit) {
	r.habboOnTiles = append(r.habboOnTiles, h)
}

func (r *RoomTile) RemoveHabboRoomUnit(h core.HabboRoomUnit) {
	for i, habbo := range r.habboOnTiles {
		if habbo == h {
			r.habboOnTiles = append(r.habboOnTiles[:i], r.habboOnTiles[i+1:]...)
		}
	}
}

// GetHeight implements core.IRoomTile.
func (r *RoomTile) GetHeight() float32 {
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

func NewRoomTile(x int32, y int32, height float32, state core.RoomTileState) core.RoomTile {
	return &RoomTile{x: x, y: y, height: height, state: state}
}
