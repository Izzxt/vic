package rooms

import (
	"context"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/database"
	room_info "github.com/Izzxt/vic/database/rooms/room_info/querier"
)

type RoomInfo struct {
	room_info.Room
	ctx context.Context
}

// GetAllowOtherPets implements core.IRoomInfo.
func (r *RoomInfo) GetAllowOtherPets() bool {
	return r.AllowOtherPets
}

// GetAllowOtherPetsEat implements core.IRoomInfo.
func (r *RoomInfo) GetAllowOtherPetsEat() bool {
	return r.AllowOtherPetsEat
}

// GetAllowWalkthrough implements core.IRoomInfo.
func (r *RoomInfo) GetAllowWalkthrough() bool {
	return r.AllowWalkthrough
}

// GetChatHearingDistance implements core.IRoomInfo.
func (r *RoomInfo) GetChatHearingDistance() int32 {
	return r.ChatHearingDistance
}

// GetChatMode implements core.IRoomInfo.
func (r *RoomInfo) GetChatMode() int32 {
	return r.ChatMode
}

// GetChatProtection implements core.IRoomInfo.
func (r *RoomInfo) GetChatProtection() int32 {
	return r.ChatProtection
}

// GetChatScrollingSpeed implements core.IRoomInfo.
func (r *RoomInfo) GetChatScrollingSpeed() int32 {
	return r.ChatScrollingSpeed
}

// GetChatWeight implements core.IRoomInfo.
func (r *RoomInfo) GetChatWeight() int32 {
	return r.ChatWeight
}

// GetDescription implements core.IRoomInfo.
func (r *RoomInfo) GetDescription() string {
	return r.Description
}

// GetFlatCategoryId implements core.IRoomInfo.
func (r *RoomInfo) GetFlatCategoryId() int32 {
	return r.FlatCategoryID
}

// GetFloorThickness implements core.IRoomInfo.
func (r *RoomInfo) GetFloorThickness() int32 {
	return r.FloorThickness
}

// GetFloorpaper implements core.IRoomInfo.
func (r *RoomInfo) GetFloorpaper() string {
	return r.Floorpaper
}

// GetId implements core.IRoomInfo.
func (r *RoomInfo) GetId() int32 {
	return r.ID
}

// GetIsForsale implements core.IRoomInfo.
func (r *RoomInfo) GetIsForsale() bool {
	return r.IsForsale
}

// GetIsPromoted implements core.IRoomInfo.
func (r *RoomInfo) GetIsPromoted() bool {
	return r.IsPromoted
}

// GetIsPublic implements core.IRoomInfo.
func (r *RoomInfo) GetIsPublic() bool {
	return r.IsPublic
}

// GetIsStaffPicked implements core.IRoomInfo.
func (r *RoomInfo) GetIsStaffPicked() bool {
	return r.IsStaffPicked
}

// GetIsWallHidden implements core.IRoomInfo.
func (r *RoomInfo) GetIsWallHidden() bool {
	return r.IsWallHidden
}

// GetIsWiredHidden implements core.IRoomInfo.
func (r *RoomInfo) GetIsWiredHidden() bool {
	return r.IsWiredHidden
}

// GetLandscape implements core.IRoomInfo.
func (r *RoomInfo) GetLandscape() string {
	return r.Landscape
}

// GetMaxUsers implements core.IRoomInfo.
func (r *RoomInfo) GetMaxUsers() int32 {
	return r.MaxUsers
}

// GetModelId implements core.IRoomInfo.
func (r *RoomInfo) GetModelId() int32 {
	return r.ModelID
}

// GetMoveDiagonal implements core.IRoomInfo.
func (r *RoomInfo) GetMoveDiagonal() bool {
	return r.MoveDiagonal
}

// GetName implements core.IRoomInfo.
func (r *RoomInfo) GetName() string {
	return r.Name
}

// GetPassword implements core.IRoomInfo.
func (r *RoomInfo) GetPassword() string {
	return r.Password
}

// GetRollerSpeed implements core.IRoomInfo.
func (r *RoomInfo) GetRollerSpeed() int32 {
	return r.RollerSpeed
}

// GetScore implements core.IRoomInfo.
func (r *RoomInfo) GetScore() int32 {
	return r.Score
}

// GetState implements core.IRoomInfo.
func (r *RoomInfo) GetState() room_info.RoomsState {
	return r.State
}

// GetTags implements core.IRoomInfo.
func (r *RoomInfo) GetTags() string {
	return r.Tags
}

// GetTradeMode implements core.IRoomInfo.
func (r *RoomInfo) GetTradeMode() int32 {
	return r.TradeMode
}

// GetUsers implements core.IRoomInfo.
func (r *RoomInfo) GetUsers() int32 {
	return r.Users
}

// GetWallHeight implements core.IRoomInfo.
func (r *RoomInfo) GetWallHeight() int32 {
	return r.WallHeight
}

// GetWallThickness implements core.IRoomInfo.
func (r *RoomInfo) GetWallThickness() int32 {
	return r.WallThickness
}

// GetWallpaper implements core.IRoomInfo.
func (r *RoomInfo) GetWallpaper() string {
	return r.Wallpaper
}

// GetWhoCanBan implements core.IRoomInfo.
func (r *RoomInfo) GetWhoCanBan() int32 {
	return r.WhoCanBan
}

// GetWhoCanKick implements core.IRoomInfo.
func (r *RoomInfo) GetWhoCanKick() int32 {
	return r.WhoCanKick
}

// GetWhoCanMute implements core.IRoomInfo.
func (r *RoomInfo) GetWhoCanMute() int32 {
	return r.WhoCanMute
}

func (r *RoomInfo) Load(id int32) core.RoomInfo {
	db := database.GetInstance().RoomInfo()
	room, err := db.GetRoomById(r.ctx, id)
	if err != nil {
		panic(err)
	}

	r.Room = room

	return r
}

func NewRoomInfo(ctx context.Context) core.RoomInfo {
	return &RoomInfo{ctx: ctx}
}
