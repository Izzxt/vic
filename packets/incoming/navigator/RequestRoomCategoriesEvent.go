package navigator

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/navigator"
)

type RequestRoomCategoriesEvent struct{}

// Execute implements core.IIncomingMessage.
func (*RequestRoomCategoriesEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	client.Send(&navigator.RoomCategoriesComposer{RoomCategory: client.Navigator().NavigatorFlatCats().GetCategories()})
}
