package habbo

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/database"
	habbo_composer "github.com/Izzxt/vic/packets/outgoing/habbo"
	"github.com/Izzxt/vic/packets/outgoing/handshake"
	"github.com/Izzxt/vic/packets/outgoing/navigator"
)

type manager struct{ ctx context.Context }

var (
	connectedClients             = make(map[int]core.Habbo)
	mu               *sync.Mutex = &sync.Mutex{}
)

func LoginHabboWithAuthTicket(ctx context.Context, authTicket string, client core.HabboClient) {
	habbo := loadHabbo(ctx, authTicket, client)

	if h, ok := connectedClients[int(habbo.HabboInfo().ID)]; ok {
		h.Client().SendAlert("You have logged in from somewhere else.")
		time.Sleep(1 * time.Second)
		h.Client().Connection().Close()
		delete(connectedClients, int(h.HabboInfo().ID))
	}
	client.SetHabbo(habbo)
	connectedClients[int(habbo.HabboInfo().ID)] = habbo

	log.Printf("Habbo logged in: %v", habbo.HabboInfo().Username)
	// TODO: update habbo last login
	// TODO: update habbo last ip
	// TODO: update machine id
	// TODO: update online status

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
