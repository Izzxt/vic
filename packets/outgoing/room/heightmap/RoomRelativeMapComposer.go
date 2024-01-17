package heightmap

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type MapComposer struct {
	Room core.Room
}

func (r MapComposer) GetId() uint16 {
	return outgoing.RelativeMapComposer
}

func (r MapComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteInt(int32(r.Room.TileMap().GetCount() / r.Room.TileMap().GetLength()))
	compose.WriteInt(int32(r.Room.TileMap().GetCount()))

	for x := 0; x < int(r.Room.TileMap().GetWidth()); x++ {
		for y := 0; y < int(r.Room.TileMap().GetLength()); y++ {
			tile := r.Room.TileMap().GetTile(int32(x), int32(y))
			if tile != nil {
				compose.WriteShort(int16(tile.GetHeight()))
			}

			customStacking := true
			if customStacking {
				compose.WriteShort(int16(tile.GetHeight() * 256))
			}
		}
	}
	return compose
}

// position := split[x][y]
// if position == 'x' {
// compose.WriteShort(-1)
// } else {
// height := int32(position) - 87*256
// compose.WriteShort(int16(height))
// }

// model := r.Rooms.RoomModel()
// var (
// 	x int
// 	y int
// )
// compose.WriteInt(int32(model.GetSizeX() / model.GetSizeY()))
// compose.WriteInt(int32(len(model.GetMap())))

// for y = 0; y < model.GetSizeY(); y++ {
// for x = 0; x < model.GetSizeX(); x++ {
// if model.GetSquareState()[x][y] == roomstate.INVALID {
// compose.WriteShort(16191)
// } else if model.GetDoorY() == y && model.GetDoorX() == x {
// compose.WriteShort(0)
// } else {
// compose.WriteShort(int16(model.GetSquareHeight()[x][y] * 256))
// }
// }
// }
