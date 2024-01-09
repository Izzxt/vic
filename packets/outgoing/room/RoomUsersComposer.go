package room

// import (
// 	"github.com/Izzxt/vic/core"
// 	"github.com/Izzxt/vic/core/rooms"
// 	"github.com/Izzxt/vic/packets/outgoing"
// )

// type UsersRoomComposer struct{ Room rooms.IRooms }

// func (c UsersRoomComposer) GetId() uint16 {
// 	return outgoing.UsersRoomComposer
// }

// func (c UsersRoomComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
// 	//this.response.appendInt(1);
// 	compose.WriteInt(1)
// 	//this.response.appendInt(this.habbo.getHabboInfo().getId());
// 	compose.WriteInt(1)
// 	//this.response.appendString(this.habbo.getHabboInfo().getUsername());
// 	compose.WriteString("Izzxt")
// 	//this.response.appendString(this.habbo.getHabboInfo().getMotto());
// 	compose.WriteString("I love Clay!")
// 	//this.response.appendString(this.habbo.getHabboInfo().getLook());
// 	compose.WriteString("hd-200-1.lg-3058-92.hr-828-1394.ch-215-110.ha-987462863-1408.ea-1402-1408.ca-1558407-1327")
// 	//this.response.appendInt(this.habbo.getRoomUnit().getId()); //Room Unit ID
// 	compose.WriteInt(1)
// 	//this.response.appendInt(this.habbo.getRoomUnit().getX());
// 	compose.WriteInt(int32(c.Room.RoomModel().GetDoorX()))
// 	//this.response.appendInt(this.habbo.getRoomUnit().getY());
// 	compose.WriteInt(int32(c.Room.RoomModel().GetDoorY()))
// 	//this.response.appendString(this.habbo.getRoomUnit().getZ() + "");
// 	compose.WriteString(string(rune(c.Room.RoomModel().GetDoorZ())) + "")
// 	//this.response.appendInt(this.habbo.getRoomUnit().getBodyRotation().getValue());
// 	compose.WriteInt(int32(c.Room.RoomModel().GetDoorRotation()))
// 	//this.response.appendInt(1);
// 	compose.WriteInt(1)
// 	//this.response.appendString(this.habbo.getHabboInfo().getGender().name().toUpperCase());
// 	compose.WriteString("M")
// 	//this.response.appendInt(this.habbo.getHabboStats().guild != 0 ? this.habbo.getHabboStats().guild : -1);
// 	compose.WriteInt(-1)
// 	//this.response.appendInt(this.habbo.getHabboStats().guild != 0 ? 1 : -1);
// 	compose.WriteInt(-1)
// 	//
// 	//String name = "";
// 	//if (this.habbo.getHabboStats().guild != 0) {
// 	//	Guild g = Emulator.getGameEnvironment().getGuildManager().getGuild(this.habbo.getHabboStats().guild);
// 	//
// 	//	if (g != null)
// 	//		name = g.getName();
// 	//}
// 	//this.response.appendString(name);
// 	compose.WriteString("")
// 	//
// 	//this.response.appendString("");
// 	compose.WriteString("")
// 	//this.response.appendInt(this.habbo.getHabboStats().getAchievementScore());
// 	compose.WriteInt(0)
// 	//this.response.appendBoolean(true);
// 	compose.WriteBool(true)
// 	return compose
// }
