package room

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type OwnerComposer struct{}

func (r OwnerComposer) GetId() uint16 {
	return outgoing.RoomOwnerComposer
}

func (r OwnerComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	return compose
}
