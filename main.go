package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Izzxt/vic/database"
	"github.com/Izzxt/vic/hotel/habbo"
	"github.com/Izzxt/vic/hotel/habboclient"
	"github.com/Izzxt/vic/hotel/navigator"
	"github.com/Izzxt/vic/hotel/rooms"
	"github.com/Izzxt/vic/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	sql, err := sql.Open("mysql", "root:root@tcp(localhost:49152)/vic?parseTime=true")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	_ = database.Init(sql)

	habbo.NewManager(ctx)
	navigator := navigator.NewNavigatorManager(ctx)
	room := rooms.NewRoomManager(ctx)
	client := habboclient.NewHabboClient(ctx, navigator, room)

	ws := server.NewWsSocket(client)
	go ws.Start()
	net := server.NewTcpSocket(client)
	go net.Start()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
