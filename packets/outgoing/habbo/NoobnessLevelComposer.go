package habbo

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type NoobnessLevelComposer struct {
	Level core.NoobnessLevel
}

const (
	OLD_NOOBNESS_LEVEL core.NoobnessLevel = iota
	NEW_NOOBNESS_LEVEL
	REAL_NOOBNESS_LEVEL
)

func (c *NoobnessLevelComposer) GetId() uint16 {
	return outgoing.NoobnessLevelComposer
}

func (c *NoobnessLevelComposer) Compose(compose core.OutgoingPacket) core.OutgoingPacket {
	compose.WriteInt(int32(core.NoobnessLevel(c.Level)))
	return compose
}
