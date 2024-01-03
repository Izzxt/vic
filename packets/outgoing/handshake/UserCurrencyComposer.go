package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type UserCurrencyComposer struct {
	headerId uint16
	// habbo    core.IHabbo
}

// GetId implements core.IOutgoingMessage.
func (c *UserCurrencyComposer) GetId() uint16 {
	return outgoing.UserCurrencyComposer
}

// Compose implements core.IOutgoingMessage.
func (c *UserCurrencyComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	return nil
}
