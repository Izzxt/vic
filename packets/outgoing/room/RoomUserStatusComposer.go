package room

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type RoomUserStatusComposer struct{}

func (r RoomUserStatusComposer) GetId() uint16 {
	return outgoing.RoomUserStatusComposer
}

func (r RoomUserStatusComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteInt(1)

	for i := 0; i < 1; i++ {
		compose.WriteInt(1)      // Virtual ID
		compose.WriteInt(0)      // X
		compose.WriteInt(0)      // Y
		compose.WriteString("0") // Z
		compose.WriteInt(0)      // RotHead
		compose.WriteInt(0)      // RotBody
		compose.WriteString("/") // Statuses
	}
	return compose
}
