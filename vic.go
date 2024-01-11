package vic

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/database"
	"github.com/Izzxt/vic/networking"
	_ "github.com/go-sql-driver/mysql"
)

type Vic struct {
	Navigator core.INavigatorManager
	Room      core.IRoomManager
}

var (
	room      core.IRoomManager
	navigator core.INavigatorManager
	net       networking.Networking
)

func (v *Vic) Init() {
	sql, err := sql.Open("mysql", "root:root@tcp(localhost:49152)/vic?parseTime=true")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	database.Init(sql)

	navigator = v.Navigator
	room = v.Room

	host := ""
	port := 2097

	net = networking.NewNetworking(context.Background(), host, port, navigator, room)
	if err := net.StartWS(); err != nil {
		fmt.Printf("Error starting websocket: %v", err)
	}
	fmt.Printf("Started websocket on %s:%d", host, port)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
