package habbo

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/habbo"
)

type RequestMeMenuSettingsEvent struct{}

func (r RequestMeMenuSettingsEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	client.Send(&habbo.MeMenuSettingsComposer{})
}
