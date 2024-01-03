package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type UserDataComposer struct {
	headerId uint16
	// habbo    core.IHabbo
}

// GetId implements core.IOutgoingMessage.
func (c *UserDataComposer) GetId() uint16 {
	return outgoing.UserDataComposer
}

// Compose implements core.IOutgoingMessage.
func (c *UserDataComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteInt(61)                                                                                             // id
	compose.WriteString("Izzat")                                                                                     // name
	compose.WriteString("hd-200-1.lg-3058-92.hr-828-1394.ch-215-110.ha-987462863-1408.ea-1402-1408.ca-1558407-1327") // look
	compose.WriteString("")                                                                                          // real name
	compose.WriteString("?")                                                                                         // motto
	compose.WriteString("")                                                                                          // gender
	compose.WriteBool(false)                                                                                         // direct mail
	compose.WriteInt(6)                                                                                              // respectReceived
	compose.WriteInt(3)                                                                                              // respectRemaining
	compose.WriteInt(3)                                                                                              // respsctPetRemaining
	compose.WriteBool(false)                                                                                         // Friends stream active
	compose.WriteString("")                                                                                          // last online?
	compose.WriteBool(false)                                                                                         // Can change name
	compose.WriteBool(false)                                                                                         // safety lock
	return compose
}
