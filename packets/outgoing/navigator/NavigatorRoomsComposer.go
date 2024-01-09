package navigator

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type NavigatorRoomsComposer struct{}

// GetId implements core.IOutgoingMessage.
func (c *NavigatorRoomsComposer) GetId() uint16 {
	return outgoing.NavigatorRoomsComposer
}

// Compose implements core.IOutgoingMessage.
func (c *NavigatorRoomsComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteInt(0)
	// compose.WriteInt(1)
	// compose.WriteInt(0)
	// compose.WriteString("")
	// compose.WriteString("Caption")
	return compose
}
