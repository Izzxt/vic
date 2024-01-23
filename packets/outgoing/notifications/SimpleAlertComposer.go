
package notifications

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type SimpleAlertComposer struct {}

func (*SimpleAlertComposer) GetId() uint16 {
	return outgoing.SimpleAlertComposer
}

func (c *SimpleAlertComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	return compose
}

