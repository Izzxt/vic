package friends

import (
	"github.com/Izzxt/vic/core"
)

type RequestFriendsEvent struct{}

// Execute implements core.IIncomingMessage.
func (*RequestFriendsEvent) Execute(client core.HabboClient, in core.IncomingPacket) {}
