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
	"github.com/Izzxt/vic/extensions"
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
	Dsn       string
}

var (
	room      core.RoomManager
	navigator core.NavigatorManager
	net       core.Networking
)

func (v *Vic) Init() {
	ctx := context.Background()

	plugin := extensions.NewPluginManager()

	// plugin.StartPlugin()

	sql, err := sql.Open("mysql", v.Dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer func() {
		sql.Close()
		fmt.Println("Shutting down databases...")

		fmt.Println("Shutting down plugins...")
		plugin.UnloadPlugin()
	}()

	database.Init(sql)

	navigator = nv.NewNavigatorManager(ctx)
	room = rooms.NewRoomManager(ctx)

	host := ""
	port := 2097

	m := messages.NewMessages()
	m.RegisterMessages()

	net = networking.NewNetworking(ctx, host, port, m, navigator, room, plugin)

	shutCtx, cancel := context.WithCancel(ctx)

	plugin.LoadPlugin()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		<-time.After(1 * time.Second)
		fmt.Println("Shutting down...")

		fmt.Println("Shutting down rooms...")
		room.Shutdown()
		fmt.Println("Shutting down networking...")
		if err := net.Shutdown(); err != nil {
			fmt.Printf("Error shutting down websocket: %v\n", err)
		}

		cancel()
	}()

	if err := net.StartWS(); err != nil {
		if websocket.IsCloseError(err, websocket.CloseGoingAway) {
			fmt.Println("Shut down.")
		} else {
			fmt.Printf("Error starting websocket: %v\n", err)
		}
	}

	<-shutCtx.Done()
}
