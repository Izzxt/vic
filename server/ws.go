package server

import (
	"github.com/Izzxt/vic/core"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     websocket.IsWebSocketUpgrade,
}

type wsSocket struct {
	habbo      core.IHabbo
	client     core.IHabboClient
	connection *websocket.Conn
}

// Close implements socket.ISocket.
func (ws *wsSocket) Close() error {
	return ws.connection.Close()
}

// Write implements socket.ISocket.
func (ws *wsSocket) Write(data []byte) error {
	return ws.connection.WriteMessage(websocket.BinaryMessage, data)
}

// Read implements socket.ISocket.
func (ws *wsSocket) Read() (int, []byte, error) {
	return ws.connection.ReadMessage()
}

// Start implements socket.ISocket.
func (ws *wsSocket) Start() error {
	return nil
}

// Shutdown implements socket.ISocket.
func (*wsSocket) Shutdown() error {
	panic("unimplemented")
}

func NewWsSocket(client core.IHabboClient) core.ISocket {
	return &wsSocket{client: client}
}
