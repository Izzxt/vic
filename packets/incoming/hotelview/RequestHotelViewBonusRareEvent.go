package hotelview

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/hotelview"
)

type RequestHotelViewBonusRareEvent struct{}

// Execute implements core.IIncomingMessage.
func (*RequestHotelViewBonusRareEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	client.Send(&hotelview.BonusRareComposer{})
}
