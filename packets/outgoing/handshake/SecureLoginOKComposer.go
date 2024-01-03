package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type SecureLoginOKComposer struct {
	headerId uint16
}

// GetId implements core.IOutgoingMessage.
func (*SecureLoginOKComposer) GetId() uint16 {
	return outgoing.SecureLoginOkComposer
}

// Compose implements core.IOutgoingMessage.
func (*SecureLoginOKComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	return compose
}
