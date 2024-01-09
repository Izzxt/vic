package tiles

import (
	"math"
	"strings"

	"github.com/Izzxt/vic/core"
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
func (r *RoomTileMap) FindPath(start core.IRoomTile, goal core.IRoomTile) []core.IRoomTile {
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

			tentativeGScore := gScore[current] + 1
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

// GetNeighbors implements core.IRoomTileMap.
func (r *RoomTileMap) GetNeighbors(current core.IRoomTile) []core.IRoomTile {
	neighbors := make([]core.IRoomTile, 0)
	x := current.GetX()
	y := current.GetY()

	if x > 0 {
		neighbors = append(neighbors, r.tiles[x-1][y])
	}

	if x < r.width-1 {
		neighbors = append(neighbors, r.tiles[x+1][y])
	}

	if y > 0 {
		neighbors = append(neighbors, r.tiles[x][y-1])
	}

	if y < r.length-1 {
		neighbors = append(neighbors, r.tiles[x][y+1])
	}

	if x > 0 && y > 0 {
		neighbors = append(neighbors, r.tiles[x-1][y-1])
	}

	if x < r.width-1 && y < r.length-1 {
		neighbors = append(neighbors, r.tiles[x+1][y+1])
	}

	if x > 0 && y < r.length-1 {
		neighbors = append(neighbors, r.tiles[x-1][y+1])
	}

	if x < r.width-1 && y > 0 {
		neighbors = append(neighbors, r.tiles[x+1][y-1])
	}
	return neighbors
}

func (r *RoomTileMap) heuristic(start, end core.IRoomTile) int32 {
	cal := math.Abs(float64(start.GetX()-end.GetX())) + math.Abs(float64(start.GetY()-end.GetY()))
	return int32(cal)
}

// ReconstructPath implements core.IRoomTileMap.
func (*RoomTileMap) ReconstructPath(cameFrom map[core.IRoomTile]core.IRoomTile, current core.IRoomTile) []core.IRoomTile {
	path := make([]core.IRoomTile, 0)
	path = append(path, current)
	for {
		current = cameFrom[current]
		if current == nil {
			break
		}
		path = append(path, current)
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

			tiles[x][y] = NewRoomTile(int32(x), int32(y), int32(tileHeight), state)
			arrayTileCount++
		}
	}
	tileMap.tiles = tiles
	tileMap.count = int32(arrayTileCount)
	tileMap.doorTile = tileMap.GetTile(model.GetX(), model.GetY())
	tileMap.doorDir = core.RoomTileDirection(model.GetDir())
	return tileMap
}
