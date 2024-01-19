package room_units

import (
	"strconv"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type RoomUnitRemoveComposer struct {
	RoomUnit core.HabboRoomUnit
}

// Compose implements core.IOutgoingMessage.
func (c *RoomUnitRemoveComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	id := strconv.Itoa(int(c.RoomUnit.ID()))
	compose.WriteString(id)
	return compose
}

func (c *RoomUnitRemoveComposer) GetId() uint16 {
	return outgoing.RoomUnitRemoveComposer
}
