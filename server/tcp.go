package server

import (
	"fmt"
	"log"
	"net"

	"github.com/Izzxt/vic/core"
)

type tcpSocket struct {
	conn   net.Conn
	habbo  core.IHabbo
	client core.IHabboClient
}

// Close implements socket.ISocket.
func (tcp *tcpSocket) Close() error {
	return tcp.conn.Close()
}

// Write implements socket.ISocket.
func (tcp *tcpSocket) Write(data []byte) error {
	_, err := tcp.conn.Write(data)
	return err
}

// Read implements socket.ISocket.
func (tcp *tcpSocket) Read() (int, []byte, error) {
	message := make([]byte, 1024)
	length, err := tcp.conn.Read(message)
	return length, message, err
}

// Start implements socket.ISocket.
func (tcp *tcpSocket) Start() error {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:30000")
	if err != nil {
		log.Fatalf("failed to resolve tcp addr: %v", err)
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	fmt.Println("Listening on port 30000")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Error accepting: %v", err)
		}
		fmt.Println("acceped")
		tcp.conn = conn
		tcp.client.SetHabbo(tcp.habbo)
		tcp.client.SetSocket(tcp)
		tcp.client.AddClient(tcp.habbo)

		go tcp.client.ReadMessage()
	}
}

// Shutdown implements socket.ISocket.
func (*tcpSocket) Shutdown() error {
	panic("unimplemented")
}

func NewTcpSocket(client core.IHabboClient) core.ISocket {
	return &tcpSocket{client: client}
}
