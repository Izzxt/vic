package core

import users "github.com/Izzxt/vic/database/users/querier"

type IHabbo interface {
	HabboInfo() users.User
	Room() IRoom
	SetRoom(IRoom)
	Client() IHabboClient
	RoomUnit() IHabboRoomUnit
	SetRoomUnit(IHabboRoomUnit)
}

type NoobnessLevel int

type IHabboRoomUnit interface {
	ID() int32
	Habbo() IHabbo
	Room() IRoom
	CurrentTile() IRoomTile
	PreviousTile() IRoomTile
	HeadRotation() RoomTileDirection
	BodyRotation() RoomTileDirection
	SetCurrentTile(IRoomTile)
	SetPreviousTile(IRoomTile)
	SetHeadRotation(RoomTileDirection)
	SetBodyRotation(RoomTileDirection)
	SetRoom(IRoom)
	SetHabbo(IHabbo)
	Dispose()
	WalkTo(IRoomTile)
	Statuses() map[HabboRoomUnitStatus]string
}

type HabboRoomUnitStatus string
