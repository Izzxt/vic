package handshake

import (
	"github.com/Izzxt/vic/core"
)

type ReleaseVersionEvent struct{}

// Execute implements core.IIncomingMessage.
func (*ReleaseVersionEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	println("ReleaseVersionEvent :", in.ReadString())
}
