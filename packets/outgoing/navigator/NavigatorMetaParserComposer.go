package navigator

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type NavigatorMetaParserComposer struct{}

// GetId implements core.IOutgoingMessage.
func (c *NavigatorMetaParserComposer) GetId() uint16 {
	return outgoing.NavigatorMetaParserComposer
}

var categories []string = []string{"official_view", "hotel_view", "roomads_view", "myworld_view"}

// Compose implements core.IOutgoingMessage.
func (c *NavigatorMetaParserComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteInt(int32(len(categories)))
	for _, category := range categories {
		compose.WriteString(category)
		compose.WriteInt(0)
	}
	return compose
}
