package core

import (
	"context"
)

type IHabboClient interface {
	GetContext() context.Context
	ReadMessage()
	Send(out IOutgoingMessage) error
	GetSocket() ISocket
	SetSocket(socket ISocket)
	SetHabbo(habbo IHabbo)
	AddClient(habbo IHabbo)
	RemoveClient(habbo IHabbo)
	GetHabbo() IHabbo
	Navigator() INavigatorManager
	Room() IRoomManager
	SendToAll(out IOutgoingMessage)
	SendToHabbos(habbos []IHabbo, out IOutgoingMessage)
	SendToRoom(room IRoom, out IOutgoingMessage)
}
