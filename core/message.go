package core

type IMessages interface {
	RegisterMessages()
	HandleMessages(client IHabboClient, packet IIncomingPacket)
	RegisterIncomingMessage(id int16, packet IIncomingMessage)
}
