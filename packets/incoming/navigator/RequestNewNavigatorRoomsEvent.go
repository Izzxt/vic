package navigator

import (
	"fmt"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/navigator"
)

type RequestNewNavigatorRoomsEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *RequestNewNavigatorRoomsEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	// if client.GetHabbo() != nil {
	// return
	//

	category := in.ReadString()
	search := in.ReadString() // search

	fmt.Println(category)
	if category == "myworld_view" {
		client.Send(
			&navigator.NewNavigatorSearchResultsComposer{
				SearchCode:  category,
				SearchQuery: search,
				NavigatorSearchResults: []navigator.NavigatorSearchResults{
					{
						Identifier: "myworld_view", PublicName: "My Rooms",
					},
				},
			},
		)
	}
}
