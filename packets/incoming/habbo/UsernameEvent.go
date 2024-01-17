package habbo

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/unknown"
)

type UsernameEvent struct{}

func (u UsernameEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	client.Send(&unknown.IgnoredUsersComposer{})
}
