package core

import (
	"bytes"

	"context"

	navigator_flat_cats "github.com/Izzxt/vic/database/navigator/navigator_flat_cats/querier"
	navigator_public_cats "github.com/Izzxt/vic/database/navigator/navigator_public_cats/querier"
	"github.com/Izzxt/vic/list"
	"github.com/gorilla/websocket"

	users "github.com/Izzxt/vic/database/users/querier"

	bubbles_chat "github.com/Izzxt/vic/database/rooms/bubbles_chat/querier"
	room_info "github.com/Izzxt/vic/database/rooms/room_info/querier"
)

type Habbo interface {
	HabboInfo() users.User
	Room() Room
	SetRoom(Room)
	Client() HabboClient
	RoomUnit() HabboRoomUnit
	SetRoomUnit(HabboRoomUnit)
	HabboStats() HabboStats
}

type HabboStats interface {
	Context() context.Context
	Load(int32) HabboStats
	UpdateBubbleChat(int32)
}

type NoobnessLevel int

type HabboRoomUnit interface {
	ID() int32
	Habbo() Habbo
	Room() Room
	CurrentTile() RoomTile
	PreviousTile() RoomTile
	HeadRotation() RoomTileDirection
	BodyRotation() RoomTileDirection
	SetCurrentTile(RoomTile)
	SetPreviousTile(RoomTile)
	SetHeadRotation(RoomTileDirection)
	SetBodyRotation(RoomTileDirection)
	SetRoom(Room)
	SetHabbo(Habbo)
	Dispose()
	WalkTo(context.Context, RoomTile, HabboClient)
	Statuses() map[HabboRoomUnitStatus]string
}

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
	SendAlert(string)
	SendMOTDMessage(...string)
}

type Networking interface {
	StartWS() error
	Shutdown() error
}

type Messages interface {
	RegisterMessages()
	HandleMessages(client HabboClient, packet IncomingPacket)
	RegisterIncomingMessage(id int16, packet IncomingMessage)
}

type NavigatorManager interface {
	NavigatorFlatCats() NavigatorFlatCats
	NavigatorPublicCats() NavigatorPublicCats
	SearchCategory(HabboClient, string) []NavigatorSearchResults
}

type NavigatorFlatCats interface {
	GetCategories() []navigator_flat_cats.NavigatorFlatCat
	GetCategory(category int32) navigator_flat_cats.NavigatorFlatCat
}

type NavigatorPublicCats interface {
	GetCategories() []navigator_public_cats.NavigatorPublicCat
	GetCategory(category int32) navigator_public_cats.NavigatorPublicCat
}

type NavigatorSearchResults struct {
	Identifier string
	PublicName string
	Rooms      []interface{}
}

type NavigatorSearchResultsSet interface {
	Compose(compose OutgoingPacket)
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

type Room interface {
	Model() RoomModel
	Info() RoomInfo
	EnterRoom(Habbo)
	LeaveRoom(Habbo, bool)
	PrepareRoom(Habbo)
	UnloadRoom()
	SuccessEnterRoom(Habbo)
	SetModel(RoomModel)
	TileMap() RoomTileMap
	GetHabbos() []Habbo
	GetHabbo(int32) Habbo
	GetHabboByName(string) Habbo
}

type RoomTileMap interface {
	GetTile(int32, int32) RoomTile
	GetNeighbors(RoomTile) []RoomTile
	GetDistance(RoomTile, RoomTile) int32
	ReconstructPath(map[RoomTile]RoomTile, RoomTile) list.List[RoomTile]
	FindPath(RoomTile, RoomTile) list.List[RoomTile]
	GetDoorTile() RoomTile
	GetDoorDirection() RoomTileDirection
	GetCount() int32
	GetWidth() int32
	GetHeight() int32
	GetLength() int32
	GetTiles() [][]RoomTile
}

type RoomTileState int
type RoomTileDirection int

type RoomTile interface {
	GetX() int32
	GetY() int32
	GetHeight() float32
	GetState() RoomTileState
	HabboOnTiles() []HabboRoomUnit
	AddHabboRoomUnit(HabboRoomUnit)
	RemoveHabboRoomUnit(HabboRoomUnit)
	RemoveHabboOnTile(Habbo)
}

type RoomManager interface {
	Model() RoomModel
	GetRoom(int32) Room
	GetRoomsByOwnerId(ownerId int32) []room_info.GetRoomsByOwnerIdRow
	GetActiveRooms() []room_info.GetActiveRoomsRow
	CreateRoom(ownerId int32, name string, description string, modelId int32, categoryId int32, maxVisitors int32, tradeType int32) Room
	Shutdown()
}

type RoomModel interface {
	Load(int32) RoomModel
	GetModelByName(string) RoomModel
	GetId() int32
	GetName() string
	GetHeightmap() string
	GetIsClub() bool
	GetIsCustom() bool
	GetX() int32
	GetY() int32
	GetDir() int32
}

type RoomInfo interface {
	UpdateOnlineCount(count int32)
	Load(int32) RoomInfo
	GetId() int32
	GetOwnerId() int32
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
	Owner() room_info.User
}

type ChatMessage interface {
	SendMessage()
}

type BubblesChat interface {
	GetBubbleChatById(int32) *bubbles_chat.BubblesChat
	GetBubbleChatByKey(string) *bubbles_chat.BubblesChat
}

type Command interface {
	Execute(HabboClient, []string)
}

type CommandManager interface {
	Get(string) Command
	Exists(string) bool
	HandleCommand(HabboClient, string)
	Register(string, Command)
	RegisterCommands()
}
