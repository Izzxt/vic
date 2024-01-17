package room

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/room/heightmap"
)

type RequestRoomHeightmapEvent struct{}

func (e RequestRoomHeightmapEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	if client.GetHabbo() == nil {
		return
	}

	if client.GetHabbo().Room() == nil {
		return
	}

	client.Send(&heightmap.MapComposer{Room: client.GetHabbo().Room()})
	client.Send(&heightmap.HeightMapComposer{Room: client.GetHabbo().Room()})

	client.GetHabbo().Room().SuccessEnterRoom(client.GetHabbo())
}
