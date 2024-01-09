package room

import (
	"strconv"
	"strings"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type RoomUnitStatusComposer struct {
	Habbos []core.IHabbo
}

func (c *RoomUnitStatusComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteInt(int32(len(c.Habbos)))

	for _, habbo := range c.Habbos {
		habbo := habbo.(core.IHabbo)
		compose.WriteInt(habbo.RoomUnit().ID())

		compose.WriteInt(habbo.RoomUnit().PreviousTile().GetX())
		compose.WriteInt(habbo.RoomUnit().PreviousTile().GetY())

		str := strconv.Itoa(int(habbo.RoomUnit().PreviousTile().GetHeight()))
		compose.WriteString(str)

		compose.WriteInt(int32(habbo.RoomUnit().HeadRotation()))
		compose.WriteInt(int32(habbo.RoomUnit().BodyRotation()))

		statuses := habbo.RoomUnit().Statuses()
		status := strings.Builder{}

		for key, value := range statuses {
			status.WriteString("/")
			status.WriteString(string(key))
			status.WriteString(" ")
			status.WriteString(value)
			status.WriteString("/")
		}

		compose.WriteString(status.String())
	}

	return compose
}

func (r *RoomUnitStatusComposer) GetId() uint16 {
	return outgoing.RoomUnitStatusComposer
}
