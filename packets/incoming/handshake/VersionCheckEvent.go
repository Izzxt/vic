package handshake

import (
	"github.com/Izzxt/vic/core"
)

type VersionCheckEvent struct{}

// Execute implements core.IIncomingMessage.
func (*VersionCheckEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	_ = in.ReadInt()    // clientId
	_ = in.ReadString() // gordonPath
	_ = in.ReadString() // externalVariables
}
