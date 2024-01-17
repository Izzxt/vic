package tracking

import (
	"github.com/Izzxt/vic/core"
)

type EventTrackerEvent struct{}

// Execute implements core.IIncomingMessage.
func (*EventTrackerEvent) Execute(client core.HabboClient, in core.IncomingPacket) {}
