package heightmap

import (
	"strings"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type HeightMapComposer struct {
	Room core.Room
}

func (r HeightMapComposer) GetId() uint16 {
	return outgoing.HeightMapComposer
}

func (r HeightMapComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteBool(true)
	compose.WriteInt(r.Room.Info().GetWallHeight())                                      // fixed wall height
	compose.WriteString(strings.ReplaceAll(r.Room.Model().GetHeightmap(), "\r\n", "\r")) // relative Room
	return compose
}
