package habboclient

import (
	"context"
	"log"

	"github.com/Izzxt/vic/codec"
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/messages"
	"github.com/gorilla/websocket"
)

type habboClient struct {
	ctx        context.Context
	conn       *websocket.Conn
	outgoing   chan core.OutgoingMessage
	done       chan struct{}
	networking core.Networking
	clients    map[*websocket.Conn]core.HabboClient
	messages   core.Messages
	habbo      core.Habbo

	navigator core.NavigatorManager
	room      core.RoomManager
}

// Navigator implements core.HabboClient.
func (h *habboClient) Navigator() core.NavigatorManager {
	return h.navigator
}

// Room implements core.HabboClient.
func (h *habboClient) Room() core.RoomManager {
	return h.room
}

// SetNavigator implements core.HabboClient.
func (h *habboClient) SetNavigator(m core.NavigatorManager) {
	h.navigator = m
}

// SetRoom implements core.HabboClient.
func (h *habboClient) SetRoom(m core.RoomManager) {
	h.room = m
}

// GetHabbo implements core.HabboClient.
func (h *habboClient) GetHabbo() core.Habbo {
	return h.habbo
}

// SetHabbo implements core.HabboClient.
func (h *habboClient) SetHabbo(habbo core.Habbo) {
	h.habbo = habbo
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
	h.clients[conn] = h
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
			incomingPacket := messages.NewIncomingPacket(header, data)
			h.messages.HandleMessages(h, incomingPacket)
		}
	}
}

func (h *habboClient) writeMessage() {
	for {
		select {
		case <-h.done:
			h.done <- struct{}{}
			return
		case out := <-h.outgoing:
			bytes := make([]byte, 6)
			outgoingPacket := messages.NewOutgoingPacket(out.GetId(), bytes)
			compose := out.Compose(outgoingPacket)
			bytes = codec.Encode(outgoingPacket.GetHeader(), compose.GetBytes())
			err := h.conn.WriteMessage(websocket.BinaryMessage, bytes)
			if err != nil {
				log.Fatalf("Error sending packet: %v", err)
			}
		}
	}
}

// Connection implements core.HabboClient.
func (h *habboClient) Connection() *websocket.Conn {
	return h.conn
}

func NewHabboClient(ctx context.Context, conn *websocket.Conn, messages core.Messages, networking core.Networking) core.HabboClient {
	if conn == nil {
		panic("conn cannot be nil")
	}

	outgoing := make(chan core.OutgoingMessage)
	done := make(chan struct{})
	clients := make(map[*websocket.Conn]core.HabboClient)

	return &habboClient{ctx: ctx, conn: conn, outgoing: outgoing, done: done, messages: messages, networking: networking, clients: clients}
}
