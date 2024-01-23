
package notifications

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type NotificationDialogComposer struct {}

func (*NotificationDialogComposer) GetId() uint16 {
	return outgoing.NotificationDialogComposer
}

func (c *NotificationDialogComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	// TODO: implement me
	return compose
}

