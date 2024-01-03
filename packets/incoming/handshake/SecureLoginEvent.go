package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/handshake"
)

type SecureLoginEvent struct{}

// Execute implements core.IIncomingMessage.
func (*SecureLoginEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	println("SecureLoginEvent :", in.ReadString())

	client.Send(&handshake.SecureLoginOKComposer{})
}
