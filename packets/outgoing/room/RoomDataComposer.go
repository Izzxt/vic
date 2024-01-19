package room

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type RoomDataComposer struct {
	Room    core.Room
	Enter   bool
	Forward bool
}

// Compose implements core.IOutgoingMessage.
func (c *RoomDataComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteBool(c.Enter)                   // isLoading
	compose.WriteInt(c.Room.Info().GetId())      // Room id
	compose.WriteString(c.Room.Info().GetName()) // Room name

	if c.Room.Info().GetIsPublic() {
		compose.WriteInt(0)     // owner id
		compose.WriteString("") // owner name
	} else {
		compose.WriteInt(c.Room.Info().Owner().ID)          // owner id
		compose.WriteString(c.Room.Info().Owner().Username) // owner name
	}

	compose.WriteInt(0) // room access type

	compose.WriteInt(c.Room.Info().GetUsers())          // current users
	compose.WriteInt(c.Room.Info().GetMaxUsers())       // users max
	compose.WriteString(c.Room.Info().GetDescription()) // description
	compose.WriteInt(c.Room.Info().GetTradeMode())      // trade settings
	compose.WriteInt(c.Room.Info().GetScore())          // score
	compose.WriteInt(2)
	compose.WriteInt(c.Room.Info().GetFlatCategoryId()) // category
	compose.WriteInt(0)                                 // tags count
	compose.WriteString("")                             // tags

	compose.WriteInt(8)

	compose.WriteBool(c.Forward)
	compose.WriteBool(c.Room.Info().GetIsStaffPicked())
	compose.WriteBool(false)
	compose.WriteBool(false)

	compose.WriteInt(c.Room.Info().GetWhoCanMute()) // who can mute
	compose.WriteInt(c.Room.Info().GetWhoCanKick()) // who can kick
	compose.WriteInt(c.Room.Info().GetWhoCanBan())  // who can ban

	compose.WriteBool(false) // TODO: room rights

	compose.WriteInt(c.Room.Info().GetChatMode())            // chat mode
	compose.WriteInt(c.Room.Info().GetChatWeight())          // chat size
	compose.WriteInt(c.Room.Info().GetChatScrollingSpeed())  // chat speed
	compose.WriteInt(c.Room.Info().GetChatHearingDistance()) // chat distance
	compose.WriteInt(c.Room.Info().GetChatProtection())      // extra flood
	return compose
}

// GetId implements core.IOutgoingMessage.
func (*RoomDataComposer) GetId() uint16 {
	return outgoing.RoomDataComposer
}
