package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Izzxt/vic/hotel/habbo"
	"github.com/Izzxt/vic/hotel/habboclient"
	"github.com/Izzxt/vic/server"
)

func main() {
	habbo := habbo.NewHabbo()
	client := habboclient.NewHabboClient()

	ws := server.NewWsSocket(habbo, client)
	go ws.Start()
	net := server.NewTcpSocket(habbo, client)
	go net.Start()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c // This will block until you manually exists with CRl-C
}
