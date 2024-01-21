package room_chat

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type RoomUnitChatShoutComposer struct {
	RoomUnit      core.HabboRoomUnit
	Message       string
	Gesture       int32
	Bubble        int32
	MessageLength int32
}

// Compose implements core.OutgoingMessage.
func (c *RoomUnitChatShoutComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteInt(c.RoomUnit.ID())
	compose.WriteString(c.Message)
	compose.WriteInt(c.Gesture)
	compose.WriteInt(c.Bubble)
	compose.WriteInt(0)
	compose.WriteInt(c.MessageLength)
	return compose

}

// GetId implements core.OutgoingMessage.
func (*RoomUnitChatShoutComposer) GetId() uint16 {
	return outgoing.RoomUnitChatShoutComposer
}
