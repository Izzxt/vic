package navigator

import (
	"strings"

	"github.com/Izzxt/vic/core"
	navigator_flat_cats "github.com/Izzxt/vic/database/navigator/navigator_flat_cats/querier"
	"github.com/Izzxt/vic/packets/outgoing"
)

type RoomCategoriesComposer struct {
	RoomCategory []navigator_flat_cats.NavigatorFlatCat
}

// Compose implements core.IOutgoingMessage.
func (c *RoomCategoriesComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteInt(int32(len(c.RoomCategory)))
	for _, category := range c.RoomCategory {
		compose.WriteInt(category.ID)
		compose.WriteString(category.Caption)
		compose.WriteBool(true)
		compose.WriteBool(false)
		compose.WriteString(category.Caption)

		if strings.HasPrefix(category.Caption, "${") {
			compose.WriteString("")
		} else {
			compose.WriteString(category.Caption)
		}

		compose.WriteBool(false)
	}

	return compose
}

// GetId implements core.IOutgoingMessage.
func (*RoomCategoriesComposer) GetId() uint16 {
	return outgoing.RoomCategoriesComposer
}
