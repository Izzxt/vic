package room

import (
	"github.com/Izzxt/vic/core"
)

type RoomUnitWalkEvent struct{}

func (e *RoomUnitWalkEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	x := in.ReadInt()
	y := in.ReadInt()

	tile := client.GetHabbo().RoomUnit().Room().TileMap().GetTile(int32(x), int32(y))
	if tile == nil {
		return
	}

	client.GetHabbo().RoomUnit().WalkTo(tile)
}
