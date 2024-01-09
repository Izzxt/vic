package room

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type UserDataComposer struct{}

func (u UserDataComposer) GetId() uint16 {
	return outgoing.RoomUserDataComposer
}

func (u UserDataComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	// room unit id
	compose.WriteInt(1)
	// user look
	compose.WriteString("hd-200-1.lg-3058-92.hr-828-1394.ch-215-110.ha-987462863-1408.ea-1402-1408.ca-1558407-1327")
	// user gender
	compose.WriteString("M")
	// user motto
	compose.WriteString("I love Clay!")
	// user score
	compose.WriteInt(0)
	return compose
}
