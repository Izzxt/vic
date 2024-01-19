package room

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/room"
)

type RequestRoomDataEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *RequestRoomDataEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	roomId := in.ReadInt()
	forward := in.ReadBool()
	enter := in.ReadBool()

	r := client.Room().GetRoom(roomId)
	if r == nil {
		return
	}

	if forward && !enter {
		enter = false
	}

	client.Send(&room.RoomDataComposer{Room: r, Enter: enter, Forward: forward})
}
