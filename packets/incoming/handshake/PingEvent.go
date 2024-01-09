package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/handshake"
)

type PingEvent struct{}

// Execute implements core.IIncomingMessage.
func (*PingEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	client.Send(&handshake.PongComposer{Id: in.ReadInt()})
}
