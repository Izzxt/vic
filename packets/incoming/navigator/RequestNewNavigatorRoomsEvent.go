package navigator

import (
	"fmt"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/navigator"
)

type RequestNewNavigatorRoomsEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *RequestNewNavigatorRoomsEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	// if client.GetHabbo() != nil {
	// return
	//

	category := in.ReadString()
	search := in.ReadString() // search

	fmt.Println(category, search)

	c := client.Navigator()

	results := c.SearchCategory(client, category)
	client.Send(
		&navigator.NewNavigatorSearchResultsComposer{
			SearchCode:             category,
			SearchQuery:            search,
			NavigatorSearchResults: results,
		},
	)
}
