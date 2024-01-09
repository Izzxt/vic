package habbo

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type MeMenuSettingsComposer struct{}

func (m MeMenuSettingsComposer) GetId() uint16 {
	return outgoing.MeMenuSettingsComposer
}

func (m MeMenuSettingsComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteInt(1)      // volume_system
	compose.WriteInt(1)      // volume_furni
	compose.WriteInt(1)      // volume_trax
	compose.WriteBool(true)  // prefer_old_chat
	compose.WriteBool(false) // block_room_invites
	compose.WriteBool(true)  // block_camera_follow
	compose.WriteInt(1)      // ui_flags
	compose.WriteInt(1)      // chat_color
	return compose
}
