package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/handshake"
)

type UniqueIdEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *UniqueIdEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	// log := e.Log
	in.ReadString()
	machineId := in.ReadString()
	// log.Info().Str("machineId", machineId).Msg("unique id")

	client.Send(&handshake.SetUniqueIdComposer{Id: machineId})
}
