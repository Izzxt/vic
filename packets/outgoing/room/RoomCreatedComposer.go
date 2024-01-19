package room

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type RoomCreatedComposer struct {
	Room core.Room
}

// GetId implements core.OutgoingMessage.
func (*RoomCreatedComposer) GetId() uint16 {
	return outgoing.RoomCreatedComposer
}

func (r *RoomCreatedComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteInt(r.Room.Info().GetId())
	compose.WriteString(r.Room.Info().GetName())
	return compose
}
