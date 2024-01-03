package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/handshake"
)

type RequestUserCreditsEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *RequestUserCreditsEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	client.Send(&handshake.UserCreditsComposer{})
}
