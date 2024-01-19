package room

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type UserDataComposer struct {
	Habbo core.Habbo
}

func (u UserDataComposer) GetId() uint16 {
	return outgoing.RoomUserDataComposer
}

func (c UserDataComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	// room unit id
	if c.Habbo.RoomUnit() != nil {
		compose.WriteInt(c.Habbo.RoomUnit().ID())
	} else {
		compose.WriteInt(-1)
	}
	// user look
	compose.WriteString(c.Habbo.HabboInfo().Look)
	// user gender
	compose.WriteString(string(c.Habbo.HabboInfo().Gender))
	// user motto
	compose.WriteString(c.Habbo.HabboInfo().Motto)
	// user score
	compose.WriteInt(0)
	return compose
}
