package habbo

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type UserCurrencyComposer struct{ Habbo core.Habbo }

// GetId implements core.IOutgoingMessage.
func (c *UserCurrencyComposer) GetId() uint16 {
	return outgoing.UserCurrencyComposer
}

// Compose implements core.IOutgoingMessage.
func (c *UserCurrencyComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	return compose
}
