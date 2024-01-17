package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type PongComposer struct{ Id int32 }

// GetId implements core.IOutgoingMessage.
func (c *PongComposer) GetId() uint16 {
	return outgoing.PongComposer
}

// Compose implements core.IOutgoingMessage.
func (c *PongComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteInt(c.Id)
	return compose
}
