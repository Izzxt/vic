package handshake

import (
	"github.com/Izzxt/vic/core"
)

type PongEvent struct{ s int }

// Execute implements core.IIncomingMessage.
func (e *PongEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	e.s = 0
}
