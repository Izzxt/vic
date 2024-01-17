package core

import (
	"bytes"

	"context"

	navigator_flat_cats "github.com/Izzxt/vic/database/navigator/navigator_flat_cats/querier"
	navigator_public_cats "github.com/Izzxt/vic/database/navigator/navigator_public_cats/querier"
	"github.com/gorilla/websocket"

	users "github.com/Izzxt/vic/database/users/querier"
)

type Habbo interface {
	HabboInfo() users.User
	Room() Room
	SetRoom(Room)
	Client() HabboClient
	RoomUnit() HabboRoomUnit
	SetRoomUnit(HabboRoomUnit)
}

type NoobnessLevel int

type HabboRoomUnit interface {
	ID() int32
	Habbo() Habbo
	Room() Room
	CurrentTile() IRoomTile
	PreviousTile() IRoomTile
	HeadRotation() RoomTileDirection
	BodyRotation() RoomTileDirection
	SetCurrentTile(IRoomTile)
	SetPreviousTile(IRoomTile)
	SetHeadRotation(RoomTileDirection)
	SetBodyRotation(RoomTileDirection)
	SetRoom(Room)
	SetHabbo(Habbo)
	Dispose()
	WalkTo(context.Context, IRoomTile, HabboClient)
	Statuses() map[HabboRoomUnitStatus]string
}

type HabboRoomUnitStatus string

type HabboClient interface {
	GetContext() context.Context
	Send(OutgoingMessage)
	Connection() *websocket.Conn
	Listen()
	AddClient(*websocket.Conn)
	SendToRoom(Room, OutgoingMessage)
	SendToHabbos([]Habbo, OutgoingMessage)
	GetHabbo() Habbo
	SetHabbo(Habbo)
	Navigator() NavigatorManager
	SetNavigator(NavigatorManager)
	Room() RoomManager
	SetRoom(RoomManager)
}

type Networking interface {
	StartWS() error
}

type Messages interface {
	RegisterMessages()
	HandleMessages(client HabboClient, packet IncomingPacket)
	RegisterIncomingMessage(id int16, packet IncomingMessage)
}

type NavigatorManager interface {
	NavigatorFlatCats() NavigatorFlatCats
	NavigatorPublicCats() NavigatorPublicCats
}

type NavigatorFlatCats interface {
	GetCategories() []navigator_flat_cats.NavigatorFlatCat
}

type NavigatorPublicCats interface {
	GetCategories() []navigator_public_cats.NavigatorPublicCat
}

type OutgoingPacket interface {
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

type OutgoingMessage interface {
	GetId() uint16
	Compose(compose OutgoingPacket) OutgoingPacket
}

type IncomingPacket interface {
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

type IncomingMessage interface {
	Execute(client HabboClient, packet IncomingPacket)
}
