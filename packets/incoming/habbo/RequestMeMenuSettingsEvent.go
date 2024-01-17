package habbo

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/habbo"
)

type RequestMeMenuSettingsEvent struct{}

func (r RequestMeMenuSettingsEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	client.Send(&habbo.MeMenuSettingsComposer{})
}
