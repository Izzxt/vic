package core

type Vic interface {
	Init()
	NavigatorManager() NavigatorManager
	RoomManager() RoomManager
}
