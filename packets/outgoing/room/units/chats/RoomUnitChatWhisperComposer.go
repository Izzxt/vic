package room_chat

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type RoomUnitChatWhisperComposer struct {
	RoomUnit      core.HabboRoomUnit
	Message       string
	Gesture       int32
	Bubble        int32
	MessageLength int32
}

// Compose implements core.OutgoingMessage.
func (c *RoomUnitChatWhisperComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteInt(c.RoomUnit.ID())
	compose.WriteString(c.Message)
	compose.WriteInt(c.Gesture)
	compose.WriteInt(c.Bubble)
	compose.WriteInt(0)
	compose.WriteInt(c.MessageLength)
	return compose

}

// GetId implements core.OutgoingMessage.
func (*RoomUnitChatWhisperComposer) GetId() uint16 {
	return outgoing.RoomUnitChatWhisperComposer
}
