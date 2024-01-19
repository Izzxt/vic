// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package navigator_flat_cats

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type RoomsState string

const (
	RoomsStateOpen      RoomsState = "open"
	RoomsStateLocked    RoomsState = "locked"
	RoomsStatePassword  RoomsState = "password"
	RoomsStateInvisible RoomsState = "invisible"
)

func (e *RoomsState) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RoomsState(s)
	case string:
		*e = RoomsState(s)
	default:
		return fmt.Errorf("unsupported scan type for RoomsState: %T", src)
	}
	return nil
}

type NullRoomsState struct {
	RoomsState RoomsState `json:"rooms_state"`
	Valid      bool       `json:"valid"` // Valid is true if RoomsState is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRoomsState) Scan(value interface{}) error {
	if value == nil {
		ns.RoomsState, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RoomsState.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRoomsState) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.RoomsState), nil
}

type UsersGender string

const (
	UsersGenderM UsersGender = "M"
	UsersGenderF UsersGender = "F"
)

func (e *UsersGender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersGender(s)
	case string:
		*e = UsersGender(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersGender: %T", src)
	}
	return nil
}

type NullUsersGender struct {
	UsersGender UsersGender `json:"users_gender"`
	Valid       bool        `json:"valid"` // Valid is true if UsersGender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersGender) Scan(value interface{}) error {
	if value == nil {
		ns.UsersGender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersGender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersGender), nil
}

type NavigatorFlatCat struct {
	ID          int32  `json:"id"`
	MinRank     int32  `json:"min_rank"`
	CaptionSave string `json:"caption_save"`
	Caption     string `json:"caption"`
	AllowTrade  bool   `json:"allow_trade"`
	MaxUsers    int32  `json:"max_users"`
	IsPublic    bool   `json:"is_public"`
	OrderNum    int32  `json:"order_num"`
}

type NavigatorPublicCat struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	HasImage bool   `json:"has_image"`
	Visible  bool   `json:"visible"`
	OrderNum int32  `json:"order_num"`
}

type Room struct {
	ID                  int32      `json:"id"`
	OwnerID             int32      `json:"owner_id"`
	Name                string     `json:"name"`
	Description         string     `json:"description"`
	ModelID             int32      `json:"model_id"`
	Password            string     `json:"password"`
	State               RoomsState `json:"state"`
	Users               int32      `json:"users"`
	MaxUsers            int32      `json:"max_users"`
	FlatCategoryID      int32      `json:"flat_category_id"`
	Score               int32      `json:"score"`
	Floorpaper          string     `json:"floorpaper"`
	Wallpaper           string     `json:"wallpaper"`
	Landscape           string     `json:"landscape"`
	WallThickness       int32      `json:"wall_thickness"`
	WallHeight          int32      `json:"wall_height"`
	FloorThickness      int32      `json:"floor_thickness"`
	Tags                string     `json:"tags"`
	IsPublic            bool       `json:"is_public"`
	IsStaffPicked       bool       `json:"is_staff_picked"`
	AllowOtherPets      bool       `json:"allow_other_pets"`
	AllowOtherPetsEat   bool       `json:"allow_other_pets_eat"`
	AllowWalkthrough    bool       `json:"allow_walkthrough"`
	IsWallHidden        bool       `json:"is_wall_hidden"`
	ChatMode            int32      `json:"chat_mode"`
	ChatWeight          int32      `json:"chat_weight"`
	ChatScrollingSpeed  int32      `json:"chat_scrolling_speed"`
	ChatHearingDistance int32      `json:"chat_hearing_distance"`
	ChatProtection      int32      `json:"chat_protection"`
	WhoCanMute          int32      `json:"who_can_mute"`
	WhoCanKick          int32      `json:"who_can_kick"`
	WhoCanBan           int32      `json:"who_can_ban"`
	RollerSpeed         int32      `json:"roller_speed"`
	IsPromoted          bool       `json:"is_promoted"`
	TradeMode           int32      `json:"trade_mode"`
	MoveDiagonal        bool       `json:"move_diagonal"`
	IsWiredHidden       bool       `json:"is_wired_hidden"`
	IsForsale           bool       `json:"is_forsale"`
}

type RoomModel struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	Heightmap string `json:"heightmap"`
	IsClub    bool   `json:"is_club"`
	IsCustom  bool   `json:"is_custom"`
	X         int32  `json:"x"`
	Y         int32  `json:"y"`
	Dir       int32  `json:"dir"`
}

type User struct {
	ID                 int32       `json:"id"`
	Username           string      `json:"username"`
	Password           string      `json:"password"`
	AuthTicket         string      `json:"auth_ticket"`
	Email              string      `json:"email"`
	RankID             int32       `json:"rank_id"`
	AccountCreatedDate time.Time   `json:"account_created_date"`
	LastOnlineDate     time.Time   `json:"last_online_date"`
	IsOnline           bool        `json:"is_online"`
	Motto              string      `json:"motto"`
	Look               string      `json:"look"`
	Gender             UsersGender `json:"gender"`
	IpRegister         string      `json:"ip_register"`
	IpCurrent          string      `json:"ip_current"`
	HomeRoom           int32       `json:"home_room"`
}
