package habbo

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/habbo"
)

type RequestUserDataEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *RequestUserDataEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	if client.GetHabbo() == nil {
		if err := client.GetSocket().Close(); err != nil {
			panic(err)
		}
		return
	}

	client.Send(&habbo.UserDataComposer{Habbo: client.GetHabbo()})
	client.Send(&habbo.UserPerksComposer{})

	client.Send(&habbo.MeMenuSettingsComposer{})
}
