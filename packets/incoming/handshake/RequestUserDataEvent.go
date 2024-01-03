package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/handshake"
)

type RequestUserDataEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *RequestUserDataEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	// if client.GetHabbo() == nil {
	// 	client.GetHabbo().GetConnection().Close()
	// 	return
	// }

	client.Send(&handshake.UserDataComposer{})
}
