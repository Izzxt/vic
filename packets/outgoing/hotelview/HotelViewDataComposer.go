package hotelview

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type HotelViewDataComposer struct {
	Data string
	Key  string
}

// Compose implements core.IOutgoingMessage.
func (c *HotelViewDataComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteString(c.Data)
	compose.WriteString(c.Key)
	return compose
}

// GetId implements core.IOutgoingMessage.
func (*HotelViewDataComposer) GetId() uint16 {
	return outgoing.HotelViewDataComposer
}
