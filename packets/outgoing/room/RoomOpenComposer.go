package room

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type OpenComposer struct{}

// Compose implements core.IOutgoingMessage.
func (*OpenComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	return compose
}

// GetId implements core.IOutgoingMessage.
func (*OpenComposer) GetId() uint16 {
	return outgoing.RoomOpenComposer
}
