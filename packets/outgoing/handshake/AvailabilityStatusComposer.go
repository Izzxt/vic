package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type AvailabilityStatusComposer struct{}

// GetId implements core.IOutgoingMessage.
func (c *AvailabilityStatusComposer) GetId() uint16 {
	return outgoing.AvailabilityStatusComposer
}

// Compose implements core.IOutgoingMessage.
func (*AvailabilityStatusComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteBool(true)
	compose.WriteBool(false)
	compose.WriteBool(true)
	return compose
}
