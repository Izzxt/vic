package core

import "bytes"

type IOutgoingPacket interface {
	Writebyte(value byte)
	WriteShort(value int16)
	WriteInt(value int32)
	WriteUInt(value uint32)
	WriteBool(value bool)
	WriteLong(value int64)
	WriteString(value string)
	WriteDouble(value []byte)
	GetHeader() uint16
	GetBytes() []byte
	GetBuffer() bytes.Buffer
}

type IOutgoingMessage interface {
	GetId() uint16
	Compose(compose IOutgoingPacket) IOutgoingPacket
}
