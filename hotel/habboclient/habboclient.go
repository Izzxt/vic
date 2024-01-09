package habboclient

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/Izzxt/vic/codec"
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/messages"
)

var (
	connectedClients = make(map[core.IHabbo]core.IHabboClient)
	clientSocket     = make(map[core.ISocket]core.IHabboClient)
)

type habboClient struct {
	ctx         context.Context
	socket      core.ISocket
	clientMutex sync.Mutex
	habbo       core.IHabbo
	navigator   core.INavigatorManager
	room        core.IRoomManager
}

// Navigator implements core.IHabboClient.
func (h *habboClient) Navigator() core.INavigatorManager {
	return h.navigator
}

func (h *habboClient) Room() core.IRoomManager {
	return h.room
}

// GetContext implements core.IHabboClient.
func (h *habboClient) GetContext() context.Context {
	return h.ctx
}

// SetHabbo implements core.IHabboClient.
func (h *habboClient) SetHabbo(habbo core.IHabbo) {
	h.habbo = habbo
}

// AddClient implements core.IHabboClient.
func (h *habboClient) AddClient(habbo core.IHabbo) {
	h.clientMutex.Lock()
	defer h.clientMutex.Unlock()
	connectedClients[habbo] = h
	clientSocket[h.GetSocket()] = h
}

func (h *habboClient) GetHabbo() core.IHabbo {
	return h.habbo
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
			h.RemoveClient(h.habbo)
			h.socket.Close()
			fmt.Printf("Error reading message: %v", err)
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
	h.clientMutex.Lock()
	defer h.clientMutex.Unlock()
	bytes := make([]byte, 6)
	outgoingPacket := messages.NewOutgoingPacket(out.GetId(), bytes)
	compose := out.Compose(outgoingPacket)
	bytes = codec.Encode(outgoingPacket.GetHeader(), compose.GetBytes())
	err := h.GetSocket().Write(bytes)
	if err != nil {
		log.Fatalf("Error sending packet: %v", err)
	}
	fmt.Printf("\033[37mSent packet: %v to %v\n\033[0m", outgoingPacket.GetHeader(), h.GetHabbo().HabboInfo().Username)
	// Write bytes to a binary file
	file, err := os.Create("logs/" + fmt.Sprintf("output-%d.bin", outgoingPacket.GetHeader()))
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
	return nil
}

// SendToAll implements core.IHabboClient.
func (h *habboClient) SendToAll(out core.IOutgoingMessage) {
	for _, client := range connectedClients {
		client.Send(out)
	}
}

func (h *habboClient) SendToHabbos(habbos []core.IHabbo, out core.IOutgoingMessage) {
	for _, habbo := range habbos {
		client := habbo.(core.IHabbo).Client()
		client.Send(out)
	}
}

func (h *habboClient) SendToRoom(room core.IRoom, out core.IOutgoingMessage) {
	for _, habbo := range room.GetHabbos() {
		habbo.Client().Send(out)
	}
}

// SetSocket implements core.IHabboClient.
func (h *habboClient) SetSocket(socket core.ISocket) {
	h.socket = socket
}

func NewHabboClient(ctx context.Context, navigator core.INavigatorManager, room core.IRoomManager) core.IHabboClient {
	return &habboClient{ctx: ctx, clientMutex: sync.Mutex{}, navigator: navigator, room: room}
}
