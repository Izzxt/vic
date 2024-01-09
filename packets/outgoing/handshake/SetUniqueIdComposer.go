package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type SetUniqueIdComposer struct{ Id string }

// GetId implements core.IOutgoingMessage.
func (c *SetUniqueIdComposer) GetId() uint16 {
	return outgoing.SetUniqueIdComposer
}

// Compose implements core.IOutgoingMessage.
func (c *SetUniqueIdComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteString(c.Id)
	return compose
}
