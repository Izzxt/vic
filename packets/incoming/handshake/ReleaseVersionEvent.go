package handshake

import (
	"github.com/Izzxt/vic/core"
)

type ReleaseVersionEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *ReleaseVersionEvent) Execute(gameClient core.IHabboClient, in core.IIncomingPacket) {
	build := in.ReadString()
	clientType := in.ReadString()
	clientPlatform := in.ReadInt()
	clientDeviceType := in.ReadInt()
	println(build, clientType, clientPlatform, clientDeviceType)
}
