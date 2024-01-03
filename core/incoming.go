package core

type IIncomingPacket interface {
	Readbyte() byte
	ReadShort() int16
	ReadUShort() uint16
	ReadInt() int32
	ReadUInt() uint32
	ReadBool() bool
	ReadString() string
	ReadBytes(length int) []byte
	GetHeader() int16
	GetBytes() []byte
}

type IIncomingMessage interface {
	Execute(client IHabboClient, packet IIncomingPacket)
}
