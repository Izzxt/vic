package room

import (
	"fmt"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/room"
)

type RoomCreateEvent struct{}

func (e *RoomCreateEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	habbo := client.GetHabbo()
	if habbo == nil {
		return
	}

	roomName := in.ReadString()
	roomDescription := in.ReadString()
	modelName := in.ReadString()
	categoryId := in.ReadInt()
	maxVisitors := in.ReadInt()
	tradeType := in.ReadInt()

	model := client.Room().Model().GetModelByName(modelName)
	if model == nil {
		fmt.Printf("failed to get room model by name: %v", modelName)
	}

	category := client.Navigator().NavigatorFlatCats().GetCategory(categoryId)
	if &category == nil {
		fmt.Printf("failed to get navigator flat category by id: %v", categoryId)
	}

	// Trade settings: 0 = no trading, 1 = owners and with rights, 2 = everyone trading

	roomInfo := client.Room().CreateRoom(habbo.HabboInfo().ID, roomName, roomDescription, model.GetId(), category.ID, maxVisitors, tradeType)

	client.Send(&room.RoomCreatedComposer{Room: roomInfo})
}
