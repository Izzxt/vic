package room

import (
	"github.com/Izzxt/vic/core"
)

type RequestRoomLoadEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *RequestRoomLoadEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	roomId := in.ReadInt()
	// pass := in.ReadString()

	room := client.Room().GetRoom(roomId)

	if room == nil {
		return
	}

	room.EnterRoom(client.GetHabbo())
}
