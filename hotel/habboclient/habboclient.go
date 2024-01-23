package habboclient

import (
	"context"
	"fmt"
	"sync"

	"github.com/Izzxt/vic/codec"
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/messages"
	"github.com/Izzxt/vic/packets/outgoing/notifications"
	"github.com/gorilla/websocket"
)

var clients map[*websocket.Conn]core.HabboClient

type habboClient struct {
	ctx  context.Context
	conn *websocket.Conn

	clientsMu sync.RWMutex

	outgoing chan core.OutgoingMessage
	done     chan struct{}

	messages   core.Messages
	networking core.Networking

	habbo   core.Habbo
	habboMu sync.RWMutex

	navigator core.NavigatorManager
	navMu     sync.RWMutex

	room   core.RoomManager
	roomMu sync.RWMutex

	inMu sync.RWMutex
}

// Navigator implements core.HabboClient.
func (h *habboClient) Navigator() core.NavigatorManager {
	h.navMu.RLock()
	defer h.navMu.RUnlock()

	return h.navigator
}

// Room implements core.HabboClient.
func (h *habboClient) Room() core.RoomManager {
	h.roomMu.RLock()
	defer h.roomMu.RUnlock()

	return h.room
}

// SetNavigator implements core.HabboClient.
func (h *habboClient) SetNavigator(m core.NavigatorManager) {
	h.navMu.Lock()
	h.navigator = m
	h.navMu.Unlock()
}

// SetRoom implements core.HabboClient.
func (h *habboClient) SetRoom(m core.RoomManager) {
	h.roomMu.Lock()
	h.room = m
	h.roomMu.Unlock()
}

// GetHabbo implements core.HabboClient.
func (h *habboClient) GetHabbo() core.Habbo {
	h.habboMu.RLock()
	defer h.habboMu.RUnlock()

	return h.habbo
}

// SetHabbo implements core.HabboClient.
func (h *habboClient) SetHabbo(habbo core.Habbo) {
	h.habboMu.Lock()
	h.habbo = habbo
	h.habboMu.Unlock()
}

// SendToHabbos implements core.HabboClient.
func (h *habboClient) SendToHabbos(habbos []core.Habbo, out core.OutgoingMessage) {
	for _, habbo := range habbos {
		habbo.Client().Send(out)
	}
}

// SendToRoom implements core.HabboClient.
func (h *habboClient) SendToRoom(room core.Room, out core.OutgoingMessage) {
	h.SendToHabbos(room.GetHabbos(), out)
}

// AddClient implements core.HabboClient.
func (h *habboClient) AddClient(conn *websocket.Conn) {
	h.clientsMu.Lock()
	clients[conn] = h
	h.clientsMu.Unlock()
}

// GetContext implements core.HabboClient.
func (h *habboClient) GetContext() context.Context {
	return h.ctx
}

// Listen implements core.HabboClient.
func (h *habboClient) Listen() {
	go h.writeMessage()
	h.readMessage()
}

// Send implements core.HabboClient.
func (h *habboClient) Send(out core.OutgoingMessage) {
	h.outgoing <- out
}

// SendAlert implements core.HabboClient.
func (h *habboClient) SendAlert(message string) {
	h.Send(&notifications.AlertComposer{Message: message})
}

// SendMOTDMessage implements core.HabboClient.
func (h *habboClient) SendMOTDMessage(messages ...string) {
	h.Send(&notifications.MOTDComposer{Messages: messages})
}

func (h *habboClient) readMessage() {
	for {
		select {
		case <-h.done:
			h.done <- struct{}{}
			return
		default:
			_, msg, err := h.conn.ReadMessage()
			if err != nil {
				h.done <- struct{}{}
				return
			}

			data, _, header := codec.Decode(msg, h)

			h.inMu.Lock()
			incomingPacket := messages.NewIncomingPacket(header, data)
			h.inMu.Unlock()

			h.inMu.RLock()
			h.messages.HandleMessages(h, incomingPacket)
			h.inMu.RUnlock()
		}
	}
}

func (h *habboClient) Dispose() {
	if h.habbo == nil {
		return
	}
	if h.habbo.Room() != nil {
		h.habbo.Room().LeaveRoom(h.habbo, true)
	}
}

func (h *habboClient) writeMessage() {
	for {
		select {
		case <-h.done:
			h.Dispose()
			return
		case out := <-h.outgoing:
			bytes := make([]byte, 6)
			outgoingPacket := messages.NewOutgoingPacket(out.GetId(), bytes)
			compose := out.Compose(outgoingPacket)
			bytes = codec.Encode(outgoingPacket.GetHeader(), compose.GetBytes())
			err := h.conn.WriteMessage(websocket.BinaryMessage, bytes)
			if err != nil {
				fmt.Printf("Error sending packet: %v", err)
			}
		}
	}
}

// Connection implements core.HabboClient.
func (h *habboClient) Connection() *websocket.Conn {
	return h.conn
}

func NewHabboClient(ctx context.Context, conn *websocket.Conn,
	messages core.Messages, networking core.Networking,
) core.HabboClient {
	// if conn == nil {
	// 	panic("conn cannot be nil")
	// }

	outgoing := make(chan core.OutgoingMessage, 100)
	done := make(chan struct{})
	clients = make(map[*websocket.Conn]core.HabboClient)

	return &habboClient{
		ctx:        ctx,
		conn:       conn,
		outgoing:   outgoing,
		done:       done,
		messages:   messages,
		networking: networking,
	}
}
