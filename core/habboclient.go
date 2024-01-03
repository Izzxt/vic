package core

type IHabboClient interface {
	ReadMessage()
	Send(out IOutgoingMessage) error
	GetSocket() ISocket
	SetSocket(socket ISocket)
	AddClient(habbo IHabbo)
	RemoveClient(habbo IHabbo)
}
