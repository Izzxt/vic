package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type PingComposer struct{}

// GetId implements core.IOutgoingMessage.
func (c *PingComposer) GetId() uint16 {
	return outgoing.PingComposer
}

// Compose implements core.IOutgoingMessage.
func (*PingComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	return compose
}
