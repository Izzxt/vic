package vic

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/database"
	nv "github.com/Izzxt/vic/hotel/navigator"
	"github.com/Izzxt/vic/hotel/rooms"
	"github.com/Izzxt/vic/messages"
	"github.com/Izzxt/vic/networking"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
)

type Vic struct {
	Navigator core.NavigatorManager
	Room      core.RoomManager
}

var (
	room      core.RoomManager
	navigator core.NavigatorManager
	net       core.Networking
)

func (v *Vic) Init() {
	ctx := context.Background()
	sql, err := sql.Open("mysql", "root:root@tcp(localhost:49152)/vic?parseTime=true")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	database.Init(sql)

	navigator = nv.NewNavigatorManager(ctx)
	room = rooms.NewRoomManager(ctx)

	host := ""
	port := 2097

	m := messages.NewMessages()
	m.RegisterMessages()

	net = networking.NewNetworking(context.Background(), host, port, m, navigator, room)

	shutCtx, cancel := context.WithCancel(ctx)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		<-time.After(1 * time.Second)
		fmt.Println("Shutting down...")

		room.Shutdown()
		if err := net.Shutdown(); err != nil {
			fmt.Printf("Error shutting down websocket: %v\n", err)
		}

		cancel()
	}()

	if err := net.StartWS(); err != nil {
		if websocket.IsCloseError(err, websocket.CloseGoingAway) {
			fmt.Println("Shutting down...")
		} else {
			fmt.Printf("Error starting websocket: %v\n", err)
		}
	}

	<-shutCtx.Done()
}
