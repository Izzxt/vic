package habbo

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type UserFigureComposer struct{ Habbo core.IHabbo }

// GetId implements core.IOutgoingMessage.
func (c *UserFigureComposer) GetId() uint16 {
	return outgoing.UserFigureComposer
}

// Compose implements core.IOutgoingMessage.
func (c *UserFigureComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteString(c.Habbo.HabboInfo().Look)
	compose.WriteString(string(c.Habbo.HabboInfo().Gender))
	return compose
}
