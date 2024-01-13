package room

import (
	"context"

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

	go client.GetHabbo().RoomUnit().WalkTo(context.Background(), tile)
}
