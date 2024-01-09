package habbo

import (
	"github.com/Izzxt/vic/core"
)

type RequestUserClubEvent struct{}

// Execute implements core.IIncomingMessage.
func (*RequestUserClubEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {}
