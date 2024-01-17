package habbo

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/database"
	habbo_composer "github.com/Izzxt/vic/packets/outgoing/habbo"
	"github.com/Izzxt/vic/packets/outgoing/handshake"
	"github.com/Izzxt/vic/packets/outgoing/navigator"
	"github.com/gorilla/websocket"
)

type manager struct{ ctx context.Context }

var (
	connectedClients             = make(map[int]core.Habbo)
	Habbos                       = make(map[*websocket.Conn]core.Habbo)
	mu               *sync.Mutex = &sync.Mutex{}
)

func LoginHabboWithAuthTicket(ctx context.Context, authTicket string, client core.HabboClient) {
	habbo := loadHabbo(ctx, authTicket, client)

	if h, ok := connectedClients[int(habbo.HabboInfo().ID)]; ok {
		fmt.Printf("Habbo already logged in: %v\n", h.HabboInfo().Username)
		delete(connectedClients, int(h.HabboInfo().ID))
		h.Client().Connection().Close()
	}
	client.SetHabbo(habbo)
	connectedClients[int(habbo.HabboInfo().ID)] = habbo

	log.Printf("Habbo logged in: %v", habbo.HabboInfo().Username)

	client.Send(&handshake.SecureLoginOKComposer{})
	client.Send(&navigator.NavigatorSettingsComposer{HomeRoomId: 0, RoomId: 0})
	client.Send(&handshake.AvailabilityStatusComposer{})
	client.Send(&habbo_composer.NoobnessLevelComposer{Level: habbo_composer.NEW_NOOBNESS_LEVEL})
	client.Send(&handshake.PingComposer{})
}

func loadHabbo(ctx context.Context, authTicket string, client core.HabboClient) core.Habbo {
	user, err := database.GetInstance().Users().GetUserByAuthTicket(context.Background(), authTicket)
	if err != nil {
		log.Fatalf("Error loading habbo: %v", err) //TODO: check for auth ticket missing
		return nil
	}
	return NewHabbo(ctx, user, client)
}

func NewManager(ctx context.Context) *manager {
	return &manager{ctx}
}
