package handshake

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/hotel/habbo"
)

type SecureLoginEvent struct{}

// Execute implements core.IIncomingMessage.
func (e *SecureLoginEvent) Execute(client core.HabboClient, in core.IncomingPacket) {
	ssoTicket := in.ReadString()
	if ssoTicket == "" {
		return
	}

	go habbo.LoginHabboWithAuthTicket(client.GetContext(), ssoTicket, client)
}
