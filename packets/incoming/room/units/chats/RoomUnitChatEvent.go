package room_chat

import (
	"github.com/Izzxt/vic/core"
	chat_message "github.com/Izzxt/vic/hotel/habbo/room/chat"
)

type RoomUnitChatEvent struct{}

// Execute implements core.IncomingMessage.
func (*RoomUnitChatEvent) Execute(client core.HabboClient, packet core.IncomingPacket) {
	message := packet.ReadString()
	styleId := packet.ReadInt()

	chatMessage := chat_message.NewChatMessage(client, packet.GetHeader(), message, styleId)
	chatMessage.SendMessage()

	client.GetHabbo().HabboStats().UpdateBubbleChat(styleId)
}
