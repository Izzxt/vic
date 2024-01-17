package room

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/room"
)

type RequestRoomDataEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *RequestRoomDataEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	// roomId := in.ReadInt()
	// forward := in.ReadBool()
	enter := in.ReadBool()

	// client.GetHabbo().EnterRoom(r)

	// e.Log.Debug().Int32("roomId", roomId).Bool("forward", forward).Bool("enter", enter).Msg("request room data event")

	client.Send(&room.RoomDataComposer{Enter: enter, Forward: true})
}
