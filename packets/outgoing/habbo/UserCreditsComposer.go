package habbo

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type UserCreditsComposer struct{ Habbo core.IHabbo }

// GetId implements core.IOutgoingMessage.
func (c *UserCreditsComposer) GetId() uint16 {
	return outgoing.UserCreditsComposer
}

// Compose implements core.IOutgoingMessage.
func (c *UserCreditsComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteString("50000.0")
	return compose
}
