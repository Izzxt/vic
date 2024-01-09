package habbo

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type UserDataComposer struct{ Habbo core.IHabbo }

// GetId implements core.IOutgoingMessage.
func (c *UserDataComposer) GetId() uint16 {
	return outgoing.UserDataComposer
}

// Compose implements core.IOutgoingMessage.
func (c *UserDataComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteInt(c.Habbo.HabboInfo().ID)
	compose.WriteString(c.Habbo.HabboInfo().Username)
	compose.WriteString(c.Habbo.HabboInfo().Look)
	compose.WriteString(string(c.Habbo.HabboInfo().Gender))
	compose.WriteString(c.Habbo.HabboInfo().Motto)
	compose.WriteString(c.Habbo.HabboInfo().Username)
	compose.WriteBool(false)
	compose.WriteInt(2000)
	compose.WriteInt(10)
	compose.WriteInt(5)
	compose.WriteBool(false)                   // Friends stream active
	compose.WriteString("01-01-1970 00:00:00") // last online?
	compose.WriteBool(false)                   // Can change name
	compose.WriteBool(false)
	return compose
}
