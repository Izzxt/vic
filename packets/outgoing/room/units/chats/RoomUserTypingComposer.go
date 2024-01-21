package room_chat

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type RoomUserTypingComposer struct {
	RoomUnit core.HabboRoomUnit
	IsTyping bool
}

func (c *RoomUserTypingComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteInt(c.RoomUnit.ID())
	if c.IsTyping {
		compose.WriteInt(1)
	} else {
		compose.WriteInt(0)
	}
	return compose
}

func (c *RoomUserTypingComposer) GetId() uint16 {
	return outgoing.RoomUserTypingComposer
}
