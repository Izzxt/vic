package habboclient

import (
	"context"
	"log"
	"sync"

	"github.com/Izzxt/vic/codec"
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/hotel/habbo"
	"github.com/Izzxt/vic/messages"
	"github.com/gorilla/websocket"
)

var (
	connectedClients = make(map[core.IHabbo]core.IHabboClient)
	clientSocket     = make(map[*websocket.Conn]core.IHabbo)
	// habbo            core.IHabbo = nil
)

type habboClient struct {
	ctx         context.Context
	socket      core.ISocket
	clientMutex *sync.Mutex
	cmu         *sync.Mutex
	navigator   core.INavigatorManager
	room        core.IRoomManager
	conn        *websocket.Conn
	habbo       core.IHabbo
}

// SetConnection implements core.IHabboClient.
func (h *habboClient) SetConnection(ws *websocket.Conn) {
	h.conn = ws
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
func (c *habboClient) SetHabbo(h core.IHabbo) {
	// c.clientMutex.Lock()
	// defer c.clientMutex.Unlock()
	habbo.Habbos[c.conn] = h
}

// Connection implements core.IHabboClient.
func (h *habboClient) Connection() *websocket.Conn {
	return h.conn
}

// AddClient implements core.IHabboClient.
func (h *habboClient) AddClient(habbo core.IHabbo) {
	h.clientMutex.Lock()
	defer h.clientMutex.Unlock()
	connectedClients[habbo] = h
	clientSocket[h.conn] = habbo
}

func (h *habboClient) GetHabbo() core.IHabbo {
	// h.clientMutex.Lock()
	// defer h.clientMutex.Unlock()
	if habbo, ok := habbo.Habbos[h.conn]; ok {
		return habbo
	}
	return nil
}

// GetSocket implements core.IHabboClient.
func (h *habboClient) GetSocket() core.ISocket {
	return h.socket
}

// ReadMessage implements core.IHabboClient.
func (h *habboClient) ReadMessage() {
}

// RemoveClient implements core.IHabboClient.
func (h *habboClient) RemoveClient(habbo core.IHabbo) {
	h.clientMutex.Lock()
	defer h.clientMutex.Unlock()
	delete(connectedClients, habbo)
}

// Send implements core.IHabboClient.
func (h *habboClient) Send(out core.IOutgoingMessage) error {
	bytes := make([]byte, 6)
	outgoingPacket := messages.NewOutgoingPacket(out.GetId(), bytes)
	compose := out.Compose(outgoingPacket)
	bytes = codec.Encode(outgoingPacket.GetHeader(), compose.GetBytes())
	err := h.conn.WriteMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		log.Fatalf("Error sending packet: %v", err)
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
	h.SendToHabbos(room.GetHabbos(), out)
}

// SetSocket implements core.IHabboClient.
func (h *habboClient) SetSocket(socket core.ISocket) {
	h.socket = socket
}

func NewHabboClient(ctx context.Context, conn *websocket.Conn, navigator core.INavigatorManager, room core.IRoomManager) core.IHabboClient {
	return &habboClient{ctx: ctx, clientMutex: &sync.Mutex{}, conn: conn, cmu: &sync.Mutex{}, navigator: navigator, room: room}
}
