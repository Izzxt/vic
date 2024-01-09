package tracking

import (
	"github.com/Izzxt/vic/core"
)

type EventTrackerEvent struct{}

// Execute implements core.IIncomingMessage.
func (*EventTrackerEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {}
