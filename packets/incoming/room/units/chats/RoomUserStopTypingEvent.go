package room_chat

import (
	"github.com/Izzxt/vic/core"
	room_chat "github.com/Izzxt/vic/packets/outgoing/room/units/chats"
)

type RoomUserStopTypingEvent struct{}

func (e *RoomUserStopTypingEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	if client.GetHabbo().Room() == nil {
		return
	}

	if client.GetHabbo().RoomUnit() == nil {
		return
	}

	client.SendToRoom(client.GetHabbo().Room(), &room_chat.RoomUserTypingComposer{RoomUnit: client.GetHabbo().RoomUnit(), IsTyping: false})
}
