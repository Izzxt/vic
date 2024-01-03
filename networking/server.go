package networking

import (
	"log"
	"net"
	"net/http"

	"github.com/Izzxt/vic/hotel/habbo"
	"github.com/Izzxt/vic/hotel/habboclient"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     websocket.IsWebSocketUpgrade,
}

type networking[T any] struct{}

// Close implements socket.ISocket.
func (*networking[T]) Close() error {
	panic("unimplemented")
}

// Read implements socket.ISocket.
func (*networking[T]) Read() (int, []byte, error) {
	panic("unimplemented")
}

// Shutdown implements socket.ISocket.
func (*networking[T]) Shutdown() error {
	panic("unimplemented")
}

func (n *networking[T]) serveWs(w http.ResponseWriter, r *http.Request) {
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Error upgrading connection: %v", err)
	}

	habbo := habbo.NewHabbo()
	client := habboclient.NewHabboClient()
	client.AddClient(habbo)

	go client.ReadMessage()
}

// Start implements socket.ISocket.
func (n *networking[T]) Start() error {
	switch n.checkType() {
	case "ws":
		http.HandleFunc("/", n.serveWs)
		log.Fatal(http.ListenAndServe(":2097", nil))

	case "tcp":
		addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:30000")
		if err != nil {
			log.Fatalf("failed to resolve tcp addr: %v", err)
		}
		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			log.Fatalf("Error listening: %v", err)
		}
		defer l.Close()

		for {
			_, err := l.Accept()
			if err != nil {
				log.Fatalf("Error accepting: %v", err)
				continue
			}
			client := habboclient.NewHabboClient()
			// tcp := server.NewTcpSocket(conn)
			// client.SetSocket(tcp)
			habbo := habbo.NewHabbo()
			client.AddClient(habbo)

			go client.ReadMessage()
		}
	}

	return nil
}

// Write implements socket.ISocket.
func (*networking[T]) Write(data []byte) error {
	panic("unimplemented")
}

func (n *networking[T]) checkType() string {
	switch any(n).(type) {
	case *networking[websocket.Conn]:
		return "ws"
	case *networking[net.Conn]:
		return "tcp"
	}
	return ""
}

func NewNetworking[T any]() *networking[T] {
	return &networking[T]{}
}
