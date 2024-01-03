package messages

import (
	"bytes"
	"encoding/binary"

	"github.com/Izzxt/vic/core"
)

type incomingPacket struct {
	header int16
	bytes  bytes.Reader
}

var data []byte

// GetBytes implements packets.IIncomingPacket.
func (in *incomingPacket) GetBytes() []byte {
	return data
}

// GetHeader implements IncomingPacket.
func (in *incomingPacket) GetHeader() int16 {
	return in.header
}

// ReadBytes implements IncomingPacket.
func (in *incomingPacket) ReadBytes(length int) []byte {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = in.Readbyte()
	}
	return bytes
}

// ReadString implements IncomingPacket.
func (i *incomingPacket) ReadString() string {
	length := i.ReadShort()
	data := i.ReadBytes(int(length))
	return string(data)
}

// ReadBool implements IncomingPacket.
func (in *incomingPacket) ReadBool() bool {
	var value bool
	binary.Read(&in.bytes, binary.BigEndian, &value)
	return value
}

// ReadInt implements IncomingPacket.
func (in *incomingPacket) ReadInt() int32 {
	var value int32
	binary.Read(&in.bytes, binary.BigEndian, &value)
	return value
}

// ReadShort implements IncomingPacket.
func (in *incomingPacket) ReadShort() int16 {
	var value int16
	binary.Read(&in.bytes, binary.BigEndian, &value)
	return value
}

// ReadUInt implements IncomingPacket.
func (in *incomingPacket) ReadUInt() uint32 {
	var value uint32
	binary.Read(&in.bytes, binary.BigEndian, &value)
	return value
}

// ReadUShort implements IncomingPacket.
func (in *incomingPacket) ReadUShort() uint16 {
	var value uint16
	binary.Read(&in.bytes, binary.BigEndian, &value)
	return value
}

// Readbyte implements IncomingPacket.
func (in *incomingPacket) Readbyte() byte {
	var value byte
	binary.Read(&in.bytes, binary.BigEndian, &value)
	return value
}

func NewIncomingPacket(header int16, b []byte) core.IIncomingPacket {
	data = b
	return &incomingPacket{header, *bytes.NewReader(b)}
}
