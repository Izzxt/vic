package handshake

import (
	"github.com/Izzxt/vic/core"
)

type PongEvent struct{ s int }

// Execute implements core.IIncomingMessage.
func (e *PongEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	e.s = 0
}
