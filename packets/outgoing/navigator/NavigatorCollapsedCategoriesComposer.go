package navigator

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type NavigatorCollapsedCategoriesComposer struct{}

// GetId implements core.IOutgoingMessage.
func (c *NavigatorCollapsedCategoriesComposer) GetId() uint16 {
	return outgoing.NavigatorCollapsedCategoriesComposer
}

// Compose implements core.IOutgoingMessage.
func (c *NavigatorCollapsedCategoriesComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteInt(0)
	return compose
}
