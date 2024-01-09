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
)

func GetRoomTileDirection(f RoomTile, t RoomTile) core.RoomTileDirection {
	if f.x == t.x && f.y == t.y {
		return DirectionNorth
	} else if f.x < t.x && f.y > t.y {
		return DirectionNorthEast
	} else if f.x == t.x && f.y < t.y {
		return DirectionEast
	} else if f.x > t.x && f.y > t.y {
		return DirectionSouthEast
	} else if f.x < t.x && f.y == t.y {
		return DirectionSouth
	} else if f.x > t.x && f.y < t.y {
		return DirectionSouthWest
	} else if f.x == t.x && f.y > t.y {
		return DirectionWest
	} else if f.x < t.x && f.y < t.y {
		return DirectionNorthWest
	}
	return DirectionNorth
}
