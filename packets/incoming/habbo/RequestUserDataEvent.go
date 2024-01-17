package habbo

import (
	"fmt"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/habbo"
)

type RequestUserDataEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *RequestUserDataEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	if client.GetHabbo() == nil {
		fmt.Printf("Error: client.GetHabbo() == nil")
	}

	client.Send(&habbo.UserDataComposer{Habbo: client.GetHabbo()})
	client.Send(&habbo.UserPerksComposer{})

	client.Send(&habbo.MeMenuSettingsComposer{})
}
