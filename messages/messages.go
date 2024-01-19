package messages

import (
	"sync"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/incoming"
	"github.com/Izzxt/vic/packets/incoming/habbo"
	"github.com/Izzxt/vic/packets/incoming/handshake"
	"github.com/Izzxt/vic/packets/incoming/navigator"
	"github.com/Izzxt/vic/packets/incoming/room"
	room_units "github.com/Izzxt/vic/packets/incoming/room/units"
)

type messages struct {
	mutex              *sync.Mutex
	handleMessageMutex *sync.RWMutex
}

var (
	incomingMessages = make(map[int16]core.IncomingMessage)
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
	// m.RegisterIncomingMessage(incoming.EventTrackerEvent, &tracking.EventTrackerEvent{}) // 4000

	// Hotel view
	// m.RegisterIncomingMessage(incoming.RequestHotelViewBonusRareEvent, &hotelview.RequestHotelViewBonusRareEvent{}) // 957
	// m.RegisterIncomingMessage(incoming.HotelViewDataEvent, &hotelview.HotelViewDataEvent{})                         // 2912

	// Friends
	// m.RegisterIncomingMessage(incoming.RequestFriendsEvent, &friends.RequestFriendsEvent{})
	// m.RegisterIncomingMessage(incoming.RequestInitFriendsEvent, &friends.RequestInitFriendsEvent{})

	// Room
	m.RegisterIncomingMessage(incoming.RequestRoomDataEvent, &room.RequestRoomDataEvent{})
	m.RegisterIncomingMessage(incoming.RequestRoomLoadEvent, &room.RequestRoomLoadEvent{})
	m.RegisterIncomingMessage(incoming.RequestHeightmapEvent, &room.RequestRoomHeightmapEvent{})
	m.RegisterIncomingMessage(incoming.RequestRoomHeightmapEvent, &room.RequestRoomHeightmapEvent{})
	m.RegisterIncomingMessage(incoming.RoomCreateEvent, &room.RoomCreateEvent{})

	// Room unit
	m.RegisterIncomingMessage(incoming.RoomUnitWalkEvent, &room_units.RoomUnitWalkEvent{})
}

// HandleMessages implements core.IMessages.
func (m *messages) HandleMessages(client core.HabboClient, packet core.IncomingPacket) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if message, ok := incomingMessages[packet.GetHeader()]; ok {
		message.Execute(client, packet)
	} else {
		// fmt.Printf("Message %d not handled\n", packet.GetHeader())
	}
}

// RegisterIncomingMessage implements core.IMessages.
func (m *messages) RegisterIncomingMessage(id int16, in core.IncomingMessage) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	incomingMessages[id] = in
}

func NewMessages() core.Messages {
	return &messages{mutex: &sync.Mutex{}, handleMessageMutex: &sync.RWMutex{}}
}
