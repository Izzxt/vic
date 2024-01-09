package navigator

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/navigator"
)

type NewNavigatorEvent struct{}

// Execute implements core.IIncomingMessage.
func (*NewNavigatorEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	client.Send(&navigator.NavigatorPrefComposer{})
	client.Send(&navigator.NavigatorMetaParserComposer{})
	client.Send(&navigator.NavigatorRoomsComposer{})
	client.Send(&navigator.NavigatorCollapsedCategoriesComposer{})
}
