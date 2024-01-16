package room_units

import (
	"strconv"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type RoomUnitComposer struct {
	Habbos []core.IHabbo
}

// TODO: implement this
func (c *RoomUnitComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteInt(int32(len(c.Habbos)))

	for _, habbo := range c.Habbos {
		compose.WriteInt(habbo.HabboInfo().ID)
		compose.WriteString(habbo.HabboInfo().Username)
		compose.WriteString(habbo.HabboInfo().Motto)
		compose.WriteString(habbo.HabboInfo().Look)

		compose.WriteInt(habbo.RoomUnit().ID())
		compose.WriteInt(habbo.RoomUnit().CurrentTile().GetX())
		compose.WriteInt(habbo.RoomUnit().CurrentTile().GetY())
		str := strconv.Itoa(int(habbo.RoomUnit().CurrentTile().GetHeight()))
		compose.WriteString(str)
		compose.WriteInt(int32(habbo.RoomUnit().BodyRotation()))

		compose.WriteInt(1)

		compose.WriteString(string(habbo.HabboInfo().Gender))

		// // TODO: guilds
		compose.WriteInt(-1)
		compose.WriteInt(-1)
		compose.WriteString("")

		compose.WriteString("")

		// TODO: achievements
		compose.WriteInt(0)

		compose.WriteBool(true)
	}
	return compose
}

func (r *RoomUnitComposer) GetId() uint16 {
	return outgoing.RoomUnitComposer
}
