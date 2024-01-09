package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type SecureLoginOKComposer struct{}

// GetId implements core.IOutgoingMessage.
func (c *SecureLoginOKComposer) GetId() uint16 {
	return outgoing.SecureLoginOKComposer
}

// Compose implements core.IOutgoingMessage.
func (*SecureLoginOKComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	return compose
}
