package messages

import (
	"bytes"
	"encoding/binary"

	"github.com/Izzxt/vic/core"
)

type outgoingPacket struct {
	header uint16
	bytes  bytes.Buffer
}

// WriteUInt implements packets.IoutgoingMessage.
func (o *outgoingPacket) WriteUInt(value uint32) {
	binary.Write(&o.bytes, binary.BigEndian, &value)
}

// GetBuffer implements packets.IoutgoingMessage.
func (o *outgoingPacket) GetBuffer() bytes.Buffer {
	return o.bytes
}

// GetHeader implements outgoingMessage.
func (o *outgoingPacket) GetHeader() uint16 {
	return o.header
}

// GetBytes implements outgoingMessage.
func (o *outgoingPacket) GetBytes() []byte {
	return o.bytes.Bytes()
}

// WriteLong implements outgoingMessage.
func (o *outgoingPacket) WriteLong(value int64) {
	binary.Write(&o.bytes, binary.BigEndian, &value)
}

// WriteBool implements outgoingMessage.
func (o *outgoingPacket) WriteBool(value bool) {
	binary.Write(&o.bytes, binary.BigEndian, &value)
}

// Writebyte implements outgoingMessage.
func (o *outgoingPacket) Writebyte(value byte) {
	binary.Write(&o.bytes, binary.BigEndian, &value)
}

// WriteDouble implements outgoingMessage.
func (*outgoingPacket) WriteDouble(value []byte) {
	panic("unimplemented")
}

// WriteInt implements outgoingMessage.
func (o *outgoingPacket) WriteInt(value int32) {
	binary.Write(&o.bytes, binary.BigEndian, &value)
}

// WriteShort implements outgoingMessage.
func (o *outgoingPacket) WriteShort(value int16) {
	binary.Write(&o.bytes, binary.BigEndian, &value)
}

// WriteString implements outgoingMessage.
func (o *outgoingPacket) WriteString(value string) {
	binary.Write(&o.bytes, binary.BigEndian, int16(len(value)))
	binary.Write(&o.bytes, binary.BigEndian, []byte(value))
}

func NewOutgoingPacket(header uint16, b []byte) core.IOutgoingPacket {
	return &outgoingPacket{
		header: header,
		bytes:  *bytes.NewBuffer(b),
	}
}
