package unknown

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type IgnoredUsersComposer struct{}

func (i IgnoredUsersComposer) GetId() uint16 {
	return outgoing.IgnoredUsersComposer
}

func (i IgnoredUsersComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteInt(0)
	return compose
}
