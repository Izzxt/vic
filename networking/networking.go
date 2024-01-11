package networking

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Izzxt/vic/codec"
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/hotel/habboclient"
	"github.com/Izzxt/vic/list"
	"github.com/Izzxt/vic/messages"
	"github.com/gorilla/websocket"
)

type Networking interface {
	GetClients() list.List[core.IHabboClient]
	StartWS() error
	StartTCP() error
	Close()
}

type networking struct {
	ctx       context.Context
	clients   list.List[core.IHabboClient]
	cmu       *sync.Mutex
	client    core.IHabboClient
	conn      *websocket.Conn
	host      string
	port      int
	wg        *sync.WaitGroup
	messages  core.IMessages
	room      core.IRoomManager
	navigator core.INavigatorManager
}

func (n *networking) addClient(client core.IHabboClient) {
	n.cmu.Lock()
	defer n.cmu.Unlock()
	if !n.clients.Contains(client) {
		n.clients.Add(client)
	}
}

func (n *networking) removeClient(client core.IHabboClient) {
	n.cmu.Lock()
	defer n.cmu.Unlock()
	if n.clients.Contains(client) {
		n.clients.Remove(client)
	}
}

func (n *networking) GetClients() list.List[core.IHabboClient] {
	return n.clients
}

func (n *networking) Close() {
	if err := n.conn.Close(); err != nil {
		fmt.Printf("Error closing connection: %v", err)
	}
}

func (n *networking) readMessage(conn *websocket.Conn) {
	n.client = habboclient.NewHabboClient(n.ctx, conn, n.navigator, n.room)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			n.removeClient(n.client)
			fmt.Printf("Error reading message: %v", err)
		}

		// fmt.Printf("Message from %s: %s\n", n.conn.RemoteAddr().String(), string(msg))
		n.conn = conn
		n.client.SetConnection(conn)
		n.addClient(n.client)

		data, _, header := codec.Decode(msg, n.client)
		incomingPacket := messages.NewIncomingPacket(header, data)
		n.messages.HandleMessages(n.client, incomingPacket)
	}
}

func (n *networking) serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Error upgrading connection: %v", err)
	}

	go n.readMessage(conn)
}

func (n *networking) StartWS() error {
	http.HandleFunc("/", n.serveWs)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", n.host, n.port), nil); err != nil {
		return err
	}
	return nil
}

func (n *networking) StartTCP() error {
	panic("not implemented")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     websocket.IsWebSocketUpgrade,
}

func NewNetworking(ctx context.Context, host string, port int, navigator core.INavigatorManager, room core.IRoomManager) Networking {
	m := messages.NewMessages()
	m.RegisterMessages()
	return &networking{clients: list.New[core.IHabboClient](0), host: host, port: port, ctx: ctx, cmu: &sync.Mutex{}, wg: &sync.WaitGroup{}, messages: m, room: room, navigator: navigator}
}
