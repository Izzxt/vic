package main

import (
	"context"

	"github.com/Izzxt/vic"
	"github.com/Izzxt/vic/hotel/navigator"
	"github.com/Izzxt/vic/hotel/rooms"
)

func main() {
	ctx := context.Background()
	vic := vic.Vic{}

	vic.Navigator = navigator.NewNavigatorManager(ctx)
	vic.Room = rooms.NewRoomManager(ctx)

	vic.Init()
}
