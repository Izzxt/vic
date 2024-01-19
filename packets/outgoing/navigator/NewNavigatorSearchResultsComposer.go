package navigator

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/incoming/navigator/category"
	"github.com/Izzxt/vic/packets/outgoing"
)

type NewNavigatorSearchResultsComposer struct {
	SearchCode             string
	SearchQuery            string
	NavigatorSearchResults []core.NavigatorSearchResults
}

// Compose implements core.IOutgoingMessage.
func (c *NewNavigatorSearchResultsComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteString(c.SearchCode)
	compose.WriteString(c.SearchQuery)

	compose.WriteInt(int32(len(c.NavigatorSearchResults)))
	for _, result := range c.NavigatorSearchResults {
		switch c.SearchCode {
		case core.NavigatorCategoryTypeMyWorld:
			category.NewNavigatorMyWorld(c.SearchCode, c.SearchQuery, result).Compose(compose)
		case core.NavigatorCategoryTypeHotel:
			category.NewNavigatorHotelView(c.SearchCode, c.SearchQuery, result).Compose(compose)
		case core.NavigatorCategoryTypeOfficial:
			// TODO: implement
		case core.NavigatorCategoryTypeRoomAds:
			// TODO: implement
		}
	}

	return compose
}

// GetId implements core.IOutgoingMessage.
func (*NewNavigatorSearchResultsComposer) GetId() uint16 {
	return outgoing.NewNavigatorSearchResultsComposer
}
