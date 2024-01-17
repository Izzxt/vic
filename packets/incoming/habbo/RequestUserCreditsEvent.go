package habbo

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/habbo"
)

type RequestUserCreditsEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *RequestUserCreditsEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	client.Send(&habbo.UserCreditsComposer{Habbo: nil})
}
