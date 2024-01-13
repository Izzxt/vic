package room

import (
	"fmt"
	"strings"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/list"
	"github.com/Izzxt/vic/packets/outgoing"
)

type RoomUnitStatusComposer struct {
	Habbos    []core.IHabbo
	RoomUnits list.List[core.IHabboRoomUnit]
}

func NewRoomUnitStatusWithHabbosComposer(habbos []core.IHabbo) *RoomUnitStatusComposer {
	return &RoomUnitStatusComposer{Habbos: habbos, RoomUnits: nil}
}

func NewRoomUnitStatusWithRoomsComposer(roomUnits core.IHabboRoomUnit) *RoomUnitStatusComposer {
	roomUnit := &RoomUnitStatusComposer{RoomUnits: list.New[core.IHabboRoomUnit](0)}
	roomUnit.RoomUnits.Add(roomUnits)
	return roomUnit
}

func (c *RoomUnitStatusComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	if c.RoomUnits != nil {
		compose.WriteInt(int32(c.RoomUnits.Len()))
		for _, room := range c.RoomUnits.Values() {
			compose.WriteInt(room.ID())

			compose.WriteInt(room.PreviousTile().GetX())
			compose.WriteInt(room.PreviousTile().GetY())

			compose.WriteString(fmt.Sprintf("%.1f", room.PreviousTile().GetHeight()))

			compose.WriteInt(int32(room.HeadRotation()))
			compose.WriteInt(int32(room.BodyRotation()))

			statuses := room.Statuses()
			status := strings.Builder{}
			status.WriteString("/flatctrl 5/")

			for key, value := range statuses {
				status.WriteString(string(key))
				status.WriteString(" ")
				status.WriteString(value)
				status.WriteString("/")
			}

			compose.WriteString(status.String())
		}
	} else if c.Habbos != nil {
		compose.WriteInt(int32(len(c.Habbos)))
		for _, habbo := range c.Habbos {
			habbo := habbo.(core.IHabbo)
			compose.WriteInt(habbo.RoomUnit().ID())

			compose.WriteInt(habbo.RoomUnit().PreviousTile().GetX())
			compose.WriteInt(habbo.RoomUnit().PreviousTile().GetY())

			compose.WriteString(fmt.Sprintf("%.1f", habbo.RoomUnit().PreviousTile().GetHeight()))

			compose.WriteInt(int32(habbo.RoomUnit().HeadRotation()))
			compose.WriteInt(int32(habbo.RoomUnit().BodyRotation()))

			statuses := habbo.RoomUnit().Statuses()
			status := strings.Builder{}
			status.WriteString("/flatctrl 5/")

			for key, value := range statuses {
				status.WriteString(string(key))
				status.WriteString(" ")
				status.WriteString(value)
				status.WriteString("/")
			}

			compose.WriteString(status.String())
		}
	}

	return compose
}

func (r *RoomUnitStatusComposer) GetId() uint16 {
	return outgoing.RoomUnitStatusComposer
}
