package habbo

import (
	"github.com/Izzxt/vic/core"
)

type RequestUserClubEvent struct{}

// Execute implements core.IIncomingMessage.
func (*RequestUserClubEvent) Execute(client core.HabboClient, in core.IncomingPacket) {}
