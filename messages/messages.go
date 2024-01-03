package messages

import (
	"sync"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/incoming"
	"github.com/Izzxt/vic/packets/incoming/handshake"
)

type messages struct {
	mutex              *sync.Mutex
	handleMessageMutex *sync.Mutex
}

var (
	incomingMessages = make(map[int16]core.IIncomingMessage)
)

// RegisterMessages implements core.IMessages.
func (m *messages) RegisterMessages() {
	// Incoming
	m.RegisterIncomingMessage(incoming.ReleaseVersionEvent, &handshake.ReleaseVersionEvent{}) // 4000
	m.RegisterIncomingMessage(incoming.SecureLoginEvent, &handshake.SecureLoginEvent{})       // 2419

	m.RegisterIncomingMessage(incoming.RequestUserDataEvent, &handshake.RequestUserDataEvent{})       // 357
	m.RegisterIncomingMessage(incoming.RequestUserCreditsEvent, &handshake.RequestUserCreditsEvent{}) // 273
}

// HandleMessages implements core.IMessages.
func (m *messages) HandleMessages(client core.IHabboClient, packet core.IIncomingPacket) {
	if message, ok := incomingMessages[packet.GetHeader()]; ok {
		m.handleMessageMutex.Lock()
		message.Execute(client, packet)
		m.handleMessageMutex.Unlock()
	}
}

// RegisterIncomingMessage implements core.IMessages.
func (m *messages) RegisterIncomingMessage(id int16, in core.IIncomingMessage) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	incomingMessages[id] = in
}

func NewMessages() core.IMessages {
	return &messages{mutex: &sync.Mutex{}, handleMessageMutex: &sync.Mutex{}}
}
