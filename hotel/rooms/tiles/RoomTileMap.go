package tiles

import (
	"container/heap"
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
	height   int32
}

// FindPath implements core.IRoomTileMap.
func (r *RoomTileMap) FindPath(start core.IRoomTile, goal core.IRoomTile) list.List[core.IRoomTile] {
	// closedSet := make(map[core.IRoomTile]bool)
	var count int = 0
	openSet := PriorityQueue{}
	heap.Init(&openSet)
	heap.Push(&openSet, &Item{value: start, index: count, priority: 0})
	cameFrom := make(map[core.IRoomTile]core.IRoomTile)

	gScore := make(map[core.IRoomTile]int32)
	gScore[start] = 0

	fScore := make(map[core.IRoomTile]int32)
	fScore[start] = r.heuristic(start, goal)

	openSetHash := list.New[core.IRoomTile](0)
	openSetHash.Add(start)

	for openSet.Len() > 0 {
		var current *Item
		// var currentScore int32 = math.MaxInt32

		current = heap.Pop(&openSet).(*Item)
		if current.value == goal {
			return r.ReconstructPath(cameFrom, current.value)
		}

		openSet.Remove(current.index)
		for _, neighbor := range r.GetNeighbors(current.value) {
			tentativeGScore := int32(getOrDefault(gScore, current.value, math.MaxInt32)) + r.getMovementCost(current.value, neighbor)
			// tentativeGScore := gScore[current.value] + r.getMovementCost(current.value, neighbor)
			if tentativeGScore < int32(getOrDefault(gScore, neighbor, math.MaxInt32)) {
				// if tentativeGScore < gScore[neighbor] {
				cameFrom[neighbor] = current.value
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + r.heuristic(neighbor, goal)

				if !openSetHash.Contains(neighbor) {
					count++
					heap.Push(&openSet, &Item{value: neighbor, priority: int(fScore[neighbor]), index: count})
					openSetHash.Add(neighbor)
				}
			}
		}
	}

	return nil
}

func getOrDefault(m map[core.IRoomTile]int32, key core.IRoomTile, defaultValue int32) int32 {
	if val, ok := m[key]; ok {
		return val
	}
	return defaultValue
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

		if adjTile.GetState() == RoomTileStateOpen && adjTile.GetHeight() <= current.GetHeight()+1 {
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
	return r.height
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
	return r.height
}

func (r *RoomTileMap) GetTile(x, y int32) core.IRoomTile {
	if x < 0 || y < 0 || x >= r.width || y >= r.height {
		return nil
	}
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

	heightmap := strings.Split(model.GetHeightmap(), "\r\n")

	tileMap.width = int32(len(heightmap[0]))
	tileMap.height = int32(len(heightmap))

	arrayTileCount := 0
	tiles := make([][]core.IRoomTile, tileMap.width)
	for x := 0; x < int(tileMap.width); x++ {
		tiles[x] = make([]core.IRoomTile, tileMap.height)
		for y := 0; y < int(tileMap.height); y++ {
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
