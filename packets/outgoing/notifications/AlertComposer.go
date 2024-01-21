package notifications

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type AlertComposer struct {
	Message string
}

func (*AlertComposer) GetId() uint16 {
	return outgoing.AlertComposer
}

func (c *AlertComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteString(c.Message)
	return compose
}
