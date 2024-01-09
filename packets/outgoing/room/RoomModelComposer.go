package room

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type ModelComposer struct {
	Room core.IRoom
}

// Compose implements core.IOutgoingMessage.
func (c *ModelComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteString(c.Room.Model().GetName())
	compose.WriteInt(c.Room.Info().GetId())
	return compose
}

// GetId implements core.IOutgoingMessage.
func (*ModelComposer) GetId() uint16 {
	return outgoing.RoomModelComposer
}
