package navigator

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type NavigatorSettingsComposer struct {
	HomeRoomId int32
	RoomId     int32
}

// GetId implements core.IOutgoingMessage.
func (c *NavigatorSettingsComposer) GetId() uint16 {
	return outgoing.NavigatorSettingsComposer
}

// Compose implements core.IOutgoingMessage.
func (c *NavigatorSettingsComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteInt(c.HomeRoomId)
	compose.WriteInt(c.RoomId)
	return compose
}
