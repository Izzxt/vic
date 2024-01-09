package heightmap

import (
	"strings"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type HeightMapComposer struct {
	Room core.IRoom
}

func (r HeightMapComposer) GetId() uint16 {
	return outgoing.HeightMapComposer
}

func (r HeightMapComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	replace := strings.ReplaceAll(r.Room.Model().GetHeightmap(), "\r", "")
	compose.WriteBool(true)
	compose.WriteInt(r.Room.Info().GetWallHeight())              // fixed wall height
	compose.WriteString(strings.ReplaceAll(replace, "\n", "\r")) // relative Room
	return compose
}
