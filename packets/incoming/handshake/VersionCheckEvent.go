package handshake

import (
	"github.com/Izzxt/vic/core"
)

type VersionCheckEvent struct{}

// Execute implements core.IIncomingMessage.
func (*VersionCheckEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	_ = in.ReadInt()    // clientId
	_ = in.ReadString() // gordonPath
	_ = in.ReadString() // externalVariables
}
