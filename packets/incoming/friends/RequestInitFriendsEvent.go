package friends

import (
	"github.com/Izzxt/vic/core"
)

type RequestInitFriendsEvent struct{}

// Execute implements core.IIncomingMessage.
func (*RequestInitFriendsEvent) Execute(client core.HabboClient, in core.IncomingPacket) {}
