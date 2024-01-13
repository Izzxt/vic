package tiles

import (
	"math"
	"strings"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/list"
)

type RoomTileMap struct {
	tiles    [][]core.IRoomTile
	doorTile core.IRoomTile
	doorDir  core.RoomTileDirection
	count    int32
	width    int32
	length   int32
}

// FindPath implements core.IRoomTileMap.
func (r *RoomTileMap) FindPath(start core.IRoomTile, goal core.IRoomTile) list.List[core.IRoomTile] {
	closedSet := make(map[core.IRoomTile]bool)
	openSet := make(map[core.IRoomTile]bool)
	openSet[start] = true
	cameFrom := make(map[core.IRoomTile]core.IRoomTile)

	gScore := make(map[core.IRoomTile]int32)
	gScore[start] = 0

	fScore := make(map[core.IRoomTile]int32)
	fScore[start] = r.heuristic(start, goal)

	for len(openSet) > 0 {
		var current core.IRoomTile
		var currentScore int32 = math.MaxInt32
		for tile := range openSet {
			if fScore[tile] < currentScore {
				current = tile
				currentScore = fScore[tile]
			}
		}

		if current == goal {
			return r.ReconstructPath(cameFrom, current)
		}

		delete(openSet, current)
		closedSet[current] = true

		for _, neighbor := range r.GetNeighbors(current) {
			if closedSet[neighbor] {
				continue
			}

			tentativeGScore := gScore[current] + r.getMovementCost(current, neighbor)
			if !openSet[neighbor] {
				openSet[neighbor] = true
			} else if tentativeGScore >= gScore[neighbor] {
				continue
			}

			cameFrom[neighbor] = current
			gScore[neighbor] = tentativeGScore
			fScore[neighbor] = tentativeGScore + r.heuristic(neighbor, goal)
		}
	}

	return nil
}

// GetDistance implements core.IRoomTileMap.
func (*RoomTileMap) GetDistance(core.IRoomTile, core.IRoomTile) int32 {
	panic("unimplemented")
}

func (r *RoomTileMap) getAdjacentTile(tile core.IRoomTile, dir core.RoomTileDirection) core.IRoomTile {
	x := tile.GetX()
	y := tile.GetY()

	switch dir {
	case DirectionNorth:
		return r.GetTile(x, y-1)
	case DirectionNorthEast:
		return r.GetTile(x+1, y-1)
	case DirectionEast:
		return r.GetTile(x+1, y)
	case DirectionSouthEast:
		return r.GetTile(x+1, y+1)
	case DirectionSouth:
		return r.GetTile(x, y+1)
	case DirectionSouthWest:
		return r.GetTile(x-1, y+1)
	case DirectionWest:
		return r.GetTile(x-1, y)
	case DirectionNorthWest:
		return r.GetTile(x-1, y-1)
	}
	return nil
}

// GetNeighbors implements core.IRoomTileMap.
func (r *RoomTileMap) GetNeighbors(current core.IRoomTile) []core.IRoomTile {
	neighbors := make([]core.IRoomTile, 0)

	for dir := core.RoomTileDirection(0); dir < DirectionLimit; dir++ {
		adjTile := r.getAdjacentTile(current, dir)
		if adjTile == nil {
			continue
		}

		if adjTile.GetState() == RoomTileStateOpen {
			neighbors = append(neighbors, adjTile)
		}
	}

	return neighbors
}

var (
	STRAIGHT_COST int32 = 10
	DIAGONAL_COST int32 = 14
)

func (r *RoomTileMap) heuristic(start, end core.IRoomTile) int32 {
	dx := int32(math.Abs(float64(start.GetX() - end.GetX())))
	dy := int32(math.Abs(float64(start.GetY() - end.GetY())))

	cal := STRAIGHT_COST*(dx+dy) + (DIAGONAL_COST-STRAIGHT_COST)*int32(math.Min(float64(dx), float64(dy)))
	return int32(cal)
}

func (r *RoomTileMap) getMovementCost(current, neighbor core.IRoomTile) int32 {
	if current.GetX() == neighbor.GetX() || current.GetY() == neighbor.GetY() {
		return STRAIGHT_COST
	}
	return DIAGONAL_COST
}

// ReconstructPath implements core.IRoomTileMap.
func (*RoomTileMap) ReconstructPath(cameFrom map[core.IRoomTile]core.IRoomTile, current core.IRoomTile) list.List[core.IRoomTile] {
	path := list.New[core.IRoomTile](0)
	path.Add(current)
	for {
		current = cameFrom[current]
		if current == nil {
			break
		}
		path.Add(current)
	}
	return path
}

// GetLength implements core.IRoomTileMap.
func (r *RoomTileMap) GetLength() int32 {
	return r.length
}

// implements getter for RoomTileMap
func (r *RoomTileMap) GetDoorTile() core.IRoomTile {
	return r.doorTile
}

// implements getter for RoomTileMap
func (r *RoomTileMap) GetDoorDirection() core.RoomTileDirection {
	return r.doorDir
}

// implements getter for RoomTileMap
func (r *RoomTileMap) GetCount() int32 {
	return r.count
}

// implements getter for RoomTileMap
func (r *RoomTileMap) GetWidth() int32 {
	return r.width
}

// implements getter for RoomTileMap
func (r *RoomTileMap) GetHeight() int32 {
	return r.length
}

func (r *RoomTileMap) GetTile(x, y int32) core.IRoomTile {
	return r.tiles[x][y]
}

func (r *RoomTileMap) GetTiles() [][]core.IRoomTile {
	return r.tiles
}

func parse(input byte) int {
	c := "0123456789abcdefghijklmnopqrstuvwxyz"
	return strings.IndexByte(c, input)
}

func NewRoomTileMap(room core.IRoom, model core.IRoomModel) core.IRoomTileMap {
	tileMap := new(RoomTileMap)

	replace := strings.ReplaceAll(model.GetHeightmap(), "\n", "")
	heightmap := strings.Split(replace, "\r")

	tileMap.width = int32(len(heightmap[0]))
	tileMap.length = int32(len(heightmap))

	arrayTileCount := 0
	tiles := make([][]core.IRoomTile, tileMap.width)
	for x := 0; x < int(tileMap.width); x++ {
		tiles[x] = make([]core.IRoomTile, tileMap.length)
		for y := 0; y < int(tileMap.length); y++ {
			heightmapChar := heightmap[y][x]
			tileHeight := parse(heightmapChar)

			var state core.RoomTileState = RoomTileStateOpen
			if tileHeight == 'x' {
				state = RoomTileStateBlocked
			}

			tiles[x][y] = NewRoomTile(int32(x), int32(y), float32(tileHeight), state)
			arrayTileCount++
		}
	}
	tileMap.tiles = tiles
	tileMap.count = int32(arrayTileCount)
	tileMap.doorTile = tileMap.GetTile(model.GetX(), model.GetY())
	tileMap.doorDir = core.RoomTileDirection(model.GetDir())
	return tileMap
}
