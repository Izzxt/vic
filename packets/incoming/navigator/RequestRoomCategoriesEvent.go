package navigator

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/navigator"
)

type RequestRoomCategoriesEvent struct{}

// Execute implements core.IIncomingMessage.
func (*RequestRoomCategoriesEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	client.Send(&navigator.RoomCategoriesComposer{RoomCategory: client.Navigator().NavigatorFlatCats().GetCategories()})
}
