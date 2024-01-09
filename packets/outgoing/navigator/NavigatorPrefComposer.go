package navigator

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type NavigatorPrefComposer struct{}

// GetId implements core.IOutgoingMessage.
func (c *NavigatorPrefComposer) GetId() uint16 {
	return outgoing.NavigatorPrefComposer
}

// Compose implements core.IOutgoingMessage.
func (c *NavigatorPrefComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteInt(68)     //X
	compose.WriteInt(42)     //Y
	compose.WriteInt(425)    //Width
	compose.WriteInt(592)    //Height
	compose.WriteBool(false) //Show or hide saved searches.
	compose.WriteInt(0)      //No idea?
	return compose
}
