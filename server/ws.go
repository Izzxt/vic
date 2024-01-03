package server

import (
	"log"
	"net/http"

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
	http.HandleFunc("/", ws.serveWs)
	log.Fatal(http.ListenAndServe(":2097", nil))
	return nil
}

// Shutdown implements socket.ISocket.
func (*wsSocket) Shutdown() error {
	panic("unimplemented")
}

func (ws *wsSocket) serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Error upgrading connection: %v", err)
	}

	ws.connection = conn
	ws.client.SetSocket(ws)
	ws.client.AddClient(ws.habbo)

	go ws.client.ReadMessage()
}

func NewWsSocket(habbo core.IHabbo, client core.IHabboClient) core.ISocket {
	return &wsSocket{habbo: habbo, client: client}
}
