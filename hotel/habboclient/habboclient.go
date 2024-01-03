package habboclient

import (
	"log"
	"sync"

	"github.com/Izzxt/vic/codec"
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/messages"
)

var (
	connectedClients = make(map[core.IHabbo]core.IHabboClient)
)

type habboClient struct {
	socket      core.ISocket
	clientMutex sync.Mutex
}

// AddClient implements core.IHabboClient.
func (h *habboClient) AddClient(habbo core.IHabbo) {
	h.clientMutex.Lock()
	defer h.clientMutex.Unlock()
	connectedClients[habbo] = h
}

// GetSocket implements core.IHabboClient.
func (h *habboClient) GetSocket() core.ISocket {
	return h.socket
}

// ReadMessage implements core.IHabboClient.
func (h *habboClient) ReadMessage() {
	m := messages.NewMessages()
	m.RegisterMessages()
	for {
		_, msg, err := h.GetSocket().Read()
		if err != nil {
			continue
		}

		data, _, header := codec.Decode(msg, h)
		incomingPacket := messages.NewIncomingPacket(header, data)
		m.HandleMessages(h, incomingPacket)
	}
}

// RemoveClient implements core.IHabboClient.
func (h *habboClient) RemoveClient(habbo core.IHabbo) {
	h.clientMutex.Lock()
	defer h.clientMutex.Unlock()
	delete(connectedClients, habbo)
}

// Send implements core.IHabboClient.
func (h *habboClient) Send(out core.IOutgoingMessage) error {
	outgoingPacket := messages.NewOutgoingPacket(out.GetId(), make([]byte, 6))
	compose := out.Compose(outgoingPacket)
	bytes := codec.Encode(outgoingPacket.GetHeader(), compose.GetBytes())
	err := h.GetSocket().Write(bytes)
	if err != nil {
		log.Fatalf("Error sending packet: %v", err)
	}
	return nil
}

// SetSocket implements core.IHabboClient.
func (h *habboClient) SetSocket(socket core.ISocket) {
	h.socket = socket
}

func NewHabboClient() core.IHabboClient {
	return &habboClient{clientMutex: sync.Mutex{}}
}
