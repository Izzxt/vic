package hotelview

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type BonusRareComposer struct{}

// Compose implements core.IOutgoingMessage.
func (*BonusRareComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteString("prizetrophy_breed_gold")
	compose.WriteInt(0)
	compose.WriteInt(120)
	compose.WriteInt(120)
	return compose
}

// GetId implements core.IOutgoingMessage.
func (*BonusRareComposer) GetId() uint16 {
	return outgoing.BonusRareComposer
}
