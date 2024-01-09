package friends

import (
	"github.com/Izzxt/vic/core"
)

type RequestInitFriendsEvent struct{}

// Execute implements core.IIncomingMessage.
func (*RequestInitFriendsEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {}
