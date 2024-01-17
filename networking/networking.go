package networking

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/hotel/habboclient"
	"github.com/gorilla/websocket"
)

type networking struct {
	ctx      context.Context
	host     string
	port     int
	messages core.Messages

	navigator core.NavigatorManager
	room      core.RoomManager
}

// Run implements Networking.
func (n *networking) StartWS() error {
	fmt.Println("Starting networking...")

	handler := func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		client := habboclient.NewHabboClient(n.ctx, conn, n.messages, n)
		client.AddClient(conn)
		client.SetNavigator(n.navigator)
		client.SetRoom(n.room)

		fmt.Println("New client connected")

		client.Listen()
	}

	http.HandleFunc("/", handler)

	return http.ListenAndServe(fmt.Sprintf("%s:%d", n.host, n.port), nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     websocket.IsWebSocketUpgrade,
}

func NewNetworking(
	ctx context.Context, host string, port int, messages core.Messages,
	navigator core.NavigatorManager,
	room core.RoomManager,
) core.Networking {
	return &networking{
		ctx:       ctx,
		host:      host,
		port:      port,
		messages:  messages,
		navigator: navigator,
		room:      room,
	}
}
