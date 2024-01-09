package habbo

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/unknown"
)

type UsernameEvent struct{}

func (u UsernameEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	client.Send(&unknown.IgnoredUsersComposer{})
}
