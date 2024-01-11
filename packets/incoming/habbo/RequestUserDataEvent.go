package habbo

import (
	"fmt"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/habbo"
)

type RequestUserDataEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *RequestUserDataEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	if client.GetHabbo() == nil {
		fmt.Printf("Error: client.GetHabbo() == nil")
	}

	client.Send(&habbo.UserDataComposer{Habbo: client.GetHabbo()})
	client.Send(&habbo.UserPerksComposer{})

	client.Send(&habbo.MeMenuSettingsComposer{})
}
