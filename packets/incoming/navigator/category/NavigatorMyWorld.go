package category

import (
	"github.com/Izzxt/vic/core"
	room_info "github.com/Izzxt/vic/database/rooms/room_info/querier"
)

type navigatorMyWorld struct {
	searchCode  string
	searchQuery string
	results     core.NavigatorSearchResults
}

func NewNavigatorMyWorld(searchCode string, searchQuery string, results core.NavigatorSearchResults) *navigatorMyWorld {
	return &navigatorMyWorld{
		searchCode:  searchCode,
		searchQuery: searchQuery,
		results:     results,
	}
}

func (c *navigatorMyWorld) Compose(compose core.OutgoingPacket) {
	compose.WriteString(c.searchCode)
	compose.WriteString(c.searchQuery)
	compose.WriteInt(int32(core.NavigatorSearchActionNone))    // action ordinal
	compose.WriteBool(bool(core.NavigatorDisplayModeExpanded)) // display mode
	compose.WriteInt(int32(core.NavigatorListModeList))        //list mode
	compose.WriteInt(int32(len(c.results.Rooms)))              //room size
	for _, r := range c.results.Rooms {
		room := r.(room_info.GetRoomsByOwnerIdRow)
		var isPublic bool = false
		var tags []string
		compose.WriteInt(room.Room.ID)      // room id
		compose.WriteString(room.Room.Name) // room name

		if isPublic {
			compose.WriteInt(0)
			compose.WriteString("")
		} else {
			compose.WriteInt(room.User.ID)          // room owner id
			compose.WriteString(room.User.Username) // room owner name
		}

		compose.WriteInt(0) // door mode

		compose.WriteInt(room.Room.Users)    // user count
		compose.WriteInt(room.Room.MaxUsers) // max user count

		compose.WriteString(room.Room.Description) // description

		compose.WriteInt(0) // trade mode
		compose.WriteInt(0) // score

		compose.WriteInt(0)
		compose.WriteInt(room.Room.FlatCategoryID) // category id

		compose.WriteInt(int32(len(tags))) // tag count
		for _, tag := range tags {
			compose.WriteString(tag)
		}

		var base = 0

		if !isPublic {
			base = base | 8
		}

		compose.WriteInt(int32(base))
	}

}
