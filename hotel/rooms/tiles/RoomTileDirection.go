package tiles

import "github.com/Izzxt/vic/core"

const (
	DirectionNorth core.RoomTileDirection = iota
	DirectionNorthEast
	DirectionEast
	DirectionSouthEast
	DirectionSouth
	DirectionSouthWest
	DirectionWest
	DirectionNorthWest
	DirectionLimit
)

func GetRoomTileDirection(f core.IRoomTile, t core.IRoomTile) core.RoomTileDirection {
	if f.GetX() == t.GetX() && f.GetY() > t.GetY() {
		return DirectionNorth
	} else if f.GetX() < t.GetX()+1 && f.GetY() > t.GetY() {
		return DirectionNorthEast
	} else if f.GetX() < t.GetX() && f.GetY() == t.GetY() {
		return DirectionEast
	} else if f.GetX() < t.GetX() {
		return DirectionSouthEast
	} else if f.GetX() == t.GetX() && f.GetY() < t.GetY() {
		return DirectionSouth
	} else if f.GetX() > t.GetX()-1 && f.GetY() < t.GetY() {
		return DirectionSouthWest
	} else if f.GetX() > t.GetX() && f.GetY() == t.GetY() {
		return DirectionWest
	} else if f.GetX() > t.GetX() {
		return DirectionNorthWest
	}
	return DirectionNorth
}
