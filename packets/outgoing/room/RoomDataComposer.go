package room

import (
	"fmt"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type RoomDataComposer struct {
	Enter   bool
	Forward bool
}

// Compose implements core.IOutgoingMessage.
func (c *RoomDataComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	fmt.Println("RoomDataComposer", c.Enter, c.Forward)
	compose.WriteBool(c.Enter)  // isLoading
	compose.WriteInt(1)         // Room id
	compose.WriteString("test") // Room name

	compose.WriteInt(1)          // owner id
	compose.WriteString("Izzxt") // owner name

	compose.WriteInt(0)           // idk
	compose.WriteInt(0)           // current users
	compose.WriteInt(10)          // users max
	compose.WriteString("a Room") // description
	compose.WriteInt(0)           // trade settings
	compose.WriteInt(1)           // score
	compose.WriteInt(0)
	compose.WriteInt(4) // category
	compose.WriteInt(0) // tags count
	//compose.WriteString("") // tags

	compose.WriteInt(8)

	compose.WriteBool(c.Forward)
	compose.WriteBool(false)
	compose.WriteBool(false)
	compose.WriteBool(false)

	compose.WriteInt(0) // who can mute
	compose.WriteInt(0) // who can kick
	compose.WriteInt(0) // who can ban

	compose.WriteBool(true)

	compose.WriteInt(0)  // chat mode
	compose.WriteInt(0)  // chat size
	compose.WriteInt(1)  // chat speed
	compose.WriteInt(50) // chat distance
	compose.WriteInt(2)  // extra flood
	return compose
}

// GetId implements core.IOutgoingMessage.
func (*RoomDataComposer) GetId() uint16 {
	return outgoing.RoomDataComposer
}
