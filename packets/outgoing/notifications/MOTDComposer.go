package notifications

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type MOTDComposer struct {
	Messages []string
}

func (*MOTDComposer) GetId() uint16 {
	return outgoing.MOTDComposer
}

func (c *MOTDComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteInt(int32(len(c.Messages)))
	for _, message := range c.Messages {
		compose.WriteString(message)
	}
	return compose
}
