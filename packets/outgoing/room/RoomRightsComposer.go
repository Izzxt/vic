package room

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type RightsComposer struct{ Level int32 }

func (r RightsComposer) GetId() uint16 {
	return outgoing.RoomRightsComposer
}

func (r RightsComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteInt(r.Level)
	return compose
}
