package core

import (
	room_info "github.com/Izzxt/vic/database/rooms/room_info/querier"
)

type IRoom interface {
	Model() IRoomModel
	Info() IRoomInfo
	EnterRoom(IHabbo)
	LeaveRoom(IHabbo, bool)
	PrepareRoom(IHabbo)
	UnloadRoom(IHabbo)
	SuccessEnterRoom(IHabbo)
	SetModel(IRoomModel)
	TileMap() IRoomTileMap
	GetHabbos() []IHabbo
}

type IRoomTileMap interface {
	GetTile(int32, int32) IRoomTile
	GetNeighbors(IRoomTile) []IRoomTile
	GetDistance(IRoomTile, IRoomTile) int32
	ReconstructPath(map[IRoomTile]IRoomTile, IRoomTile) []IRoomTile
	FindPath(IRoomTile, IRoomTile) []IRoomTile
	GetDoorTile() IRoomTile
	GetDoorDirection() RoomTileDirection
	GetCount() int32
	GetWidth() int32
	GetHeight() int32
	GetLength() int32
	GetTiles() [][]IRoomTile
}

type RoomTileState int
type RoomTileDirection int

type IRoomTile interface {
	GetX() int32
	GetY() int32
	GetHeight() int32
	GetState() RoomTileState
	HabboOnTiles() []IHabboRoomUnit
	AddHabboRoomUnit(IHabboRoomUnit)
	RemoveHabboRoomUnit(IHabboRoomUnit)
}

type IRoomManager interface {
	GetRoom(int32) IRoom
}

type IRoomModel interface {
	Load(int32) IRoomModel
	GetId() int32
	GetName() string
	GetHeightmap() string
	GetIsClub() bool
	GetIsCustom() bool
	GetX() int32
	GetY() int32
	GetDir() int32
}

type IRoomInfo interface {
	Load(int32) IRoomInfo
	GetId() int32
	GetName() string
	GetDescription() string
	GetModelId() int32
	GetPassword() string
	GetState() room_info.RoomsState
	GetUsers() int32
	GetMaxUsers() int32
	GetFlatCategoryId() int32
	GetScore() int32
	GetFloorpaper() string
	GetWallpaper() string
	GetLandscape() string
	GetWallThickness() int32
	GetWallHeight() int32
	GetFloorThickness() int32
	GetTags() string
	GetIsPublic() bool
	GetIsStaffPicked() bool
	GetAllowOtherPets() bool
	GetAllowOtherPetsEat() bool
	GetAllowWalkthrough() bool
	GetIsWallHidden() bool
	GetChatMode() int32
	GetChatWeight() int32
	GetChatScrollingSpeed() int32
	GetChatHearingDistance() int32
	GetChatProtection() int32
	GetWhoCanMute() int32
	GetWhoCanKick() int32
	GetWhoCanBan() int32
	GetRollerSpeed() int32
	GetIsPromoted() bool
	GetTradeMode() int32
	GetMoveDiagonal() bool
	GetIsWiredHidden() bool
	GetIsForsale() bool
}
