package messages

import (
	"fmt"
	"sync"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/incoming"
	"github.com/Izzxt/vic/packets/incoming/friends"
	"github.com/Izzxt/vic/packets/incoming/habbo"
	"github.com/Izzxt/vic/packets/incoming/handshake"
	"github.com/Izzxt/vic/packets/incoming/hotelview"
	"github.com/Izzxt/vic/packets/incoming/navigator"
	"github.com/Izzxt/vic/packets/incoming/room"
	room_units "github.com/Izzxt/vic/packets/incoming/room/units"
	"github.com/Izzxt/vic/packets/incoming/tracking"
)

type messages struct {
	mutex              *sync.Mutex
	handleMessageMutex *sync.Mutex
}

var (
	incomingMessages = make(map[int16]core.IIncomingMessage)
)

// RegisterMessages implements core.IMessages.
func (m *messages) RegisterMessages() {
	// Handshake
	m.RegisterIncomingMessage(incoming.ReleaseVersionEvent, &handshake.ReleaseVersionEvent{}) // 4000
	m.RegisterIncomingMessage(incoming.SecureLoginEvent, &handshake.SecureLoginEvent{})       // 2419
	m.RegisterIncomingMessage(incoming.UniqueIdEvent, &handshake.UniqueIdEvent{})             // 3521
	m.RegisterIncomingMessage(incoming.VersionCheckEvent, &handshake.VersionCheckEvent{})     // 1220
	m.RegisterIncomingMessage(incoming.PingEvent, &handshake.PingEvent{})                     // 2419
	m.RegisterIncomingMessage(incoming.PongEvent, &handshake.PongEvent{})                     // 2596

	// Users
	m.RegisterIncomingMessage(incoming.RequestUserDataEvent, &habbo.RequestUserDataEvent{})             // 357
	m.RegisterIncomingMessage(incoming.RequestUserCreditsEvent, &habbo.RequestUserCreditsEvent{})       // 273
	m.RegisterIncomingMessage(incoming.RequestUserClubEvent, &habbo.RequestUserClubEvent{})             // 273
	m.RegisterIncomingMessage(incoming.RequestMeMenuSettingsEvent, &habbo.RequestMeMenuSettingsEvent{}) // 273
	m.RegisterIncomingMessage(incoming.UsernameEvent, &habbo.UsernameEvent{})                           // 273
	m.RegisterIncomingMessage(incoming.UserFigure, &habbo.UserFigureEvent{})                            // 273

	// Navigator
	m.RegisterIncomingMessage(incoming.NewNavigatorEvent, &navigator.NewNavigatorEvent{}) // 3375
	m.RegisterIncomingMessage(incoming.RequestNewNavigatorRoomsEvent, &navigator.RequestNewNavigatorRoomsEvent{})
	m.RegisterIncomingMessage(incoming.RequestRoomCategoriesEvent, &navigator.RequestRoomCategoriesEvent{})

	// Tracking
	m.RegisterIncomingMessage(incoming.EventTrackerEvent, &tracking.EventTrackerEvent{}) // 4000

	// Hotel view
	m.RegisterIncomingMessage(incoming.RequestHotelViewBonusRareEvent, &hotelview.RequestHotelViewBonusRareEvent{}) // 957
	m.RegisterIncomingMessage(incoming.HotelViewDataEvent, &hotelview.HotelViewDataEvent{})                         // 2912

	// Friends
	m.RegisterIncomingMessage(incoming.RequestFriendsEvent, &friends.RequestFriendsEvent{})
	m.RegisterIncomingMessage(incoming.RequestInitFriendsEvent, &friends.RequestInitFriendsEvent{})

	// Room
	m.RegisterIncomingMessage(incoming.RequestRoomDataEvent, &room.RequestRoomDataEvent{})
	m.RegisterIncomingMessage(incoming.RequestRoomLoadEvent, &room.RequestRoomLoadEvent{})
	m.RegisterIncomingMessage(incoming.RequestRoomHeightmapEvent, &room.RequestRoomHeightmapEvent{})

	// Room unit
	m.RegisterIncomingMessage(incoming.RoomUnitWalkEvent, &room_units.RoomUnitWalkEvent{})
}

// HandleMessages implements core.IMessages.
func (m *messages) HandleMessages(client core.IHabboClient, packet core.IIncomingPacket) {
	if message, ok := incomingMessages[packet.GetHeader()]; ok {
		m.handleMessageMutex.Lock()
		message.Execute(client, packet)
		fmt.Printf("Message %d handled\n", packet.GetHeader())
		m.handleMessageMutex.Unlock()
	} else {
		fmt.Printf("\033[31mMessage %d not handled\n", packet.GetHeader())
	}
}

// RegisterIncomingMessage implements core.IMessages.
func (m *messages) RegisterIncomingMessage(id int16, in core.IIncomingMessage) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	incomingMessages[id] = in
}

func NewMessages() core.IMessages {
	return &messages{mutex: &sync.Mutex{}, handleMessageMutex: &sync.Mutex{}}
}
