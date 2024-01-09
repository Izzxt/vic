package habbo

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type UserPerksComposer struct{}

// Compose implements core.IOutgoingMessage.
func (*UserPerksComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteInt(15) // Count
	compose.WriteString("USE_GUIDE_TOOL")
	compose.WriteString("requirement.unfulfilled.helper_level_4") // Not required for Nitro ?
	compose.WriteBool(true)

	compose.WriteString("GIVE_GUIDE_TOURS")
	compose.WriteString("")
	compose.WriteBool(true)

	compose.WriteString("JUDGE_CHAT_REVIEWS")
	compose.WriteString("requirement.unfulfilled.helper_level_6") // Not required for Nitro ?
	compose.WriteBool(true)

	compose.WriteString("VOTE_IN_COMPETITIONS")
	compose.WriteString("requirement.unfulfilled.helper_level_2") // Not required for Nitro ?
	compose.WriteBool(true)

	compose.WriteString("CALL_ON_HELPERS")
	compose.WriteString("")
	compose.WriteBool(true)

	compose.WriteString("CITIZEN")
	compose.WriteString("")
	compose.WriteBool(true)

	compose.WriteString("TRADE")
	compose.WriteString("requirement.unfulfilled.no_trade_lock") // Not required for Nitro ?
	compose.WriteBool(true)

	compose.WriteString("HEIGHTMAP_EDITOR_BETA")
	compose.WriteString("requirement.unfulfilled.feature_disabled") // Not required for Nitro ?
	compose.WriteBool(true)

	compose.WriteString("BUILDER_AT_WORK")
	compose.WriteString("")
	compose.WriteBool(true)

	compose.WriteString("CALL_ON_HELPERS")
	compose.WriteString("")
	compose.WriteBool(true)

	compose.WriteString("CAMERA")
	compose.WriteString("")
	compose.WriteBool(true)

	compose.WriteString("NAVIGATOR_PHASE_TWO_2014")
	compose.WriteString("")
	compose.WriteBool(true)

	compose.WriteString("MOUSE_ZOOM")
	compose.WriteString("")
	compose.WriteBool(true)

	compose.WriteString("NAVIGATOR_ROOM_THUMBNAIL_CAMERA")
	compose.WriteString("")
	compose.WriteBool(true)

	compose.WriteString("HABBO_CLUB_OFFER_BETA")
	compose.WriteString("")
	compose.WriteBool(true)

	return compose
}

// GetId implements core.IOutgoingMessage.
func (*UserPerksComposer) GetId() uint16 {
	return outgoing.UserPerksComposer
}
