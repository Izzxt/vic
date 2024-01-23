package commands

import (
	"strings"

	"github.com/Izzxt/vic/core"
)

type CmdCommand struct{}

func (c *CmdCommand) Execute(client core.HabboClient, args []string) {
	msg := "<b>Commands:</b> :about, :commands, :pickall, :sit, :stand, :lay, :dance, :wave, :idle, :walk, :mute, :unmute, :kick, :ban, :unban, :alert, :teleport, :teleportto, :teleportuser, :teleportuserhere, :teleportuserto, :teleportusertohere, :roominfo, :roomusers, :roomuser, :roomuserbyid, :roomuserbyname, :roomuserbyip, :roomuserbymachineid, :roomuserbyauth, :roomuserbyrank, :roomuserbycredits, :roomuserbyduckets, :roomuserbydiamonds, :roomuserbygotw, :roomuserbyonline, :roomuserbyoffline, :roomuserbybanned, :roomuserbymuted, :roomuserbyipmuted, :roomuserbymachineidmuted, :roomuserbyauthmuted, :roomuserbyrankmuted, :roomuserbycreditsmuted, :roomuserbyducketsmuted, :roomuserbydiamondsmuted, :roomuserbygotwmuted, :roomuserbyonlinemuted, :roomuserbyofflinemuted, :roomuserbybannedmuted, :roomuserbyipmutedmuted, :roomuserbymachineidmutedmuted, :roomuserbyauthmutedmuted, :roomuserbyrankmutedmuted, :roomuserbycreditsmutedmuted, :roomuserbyducketsmutedmuted, :roomuserbydiamondsmutedmuted, :roomuserbygotwmutedmuted, :roomuserbyonlinemutedmuted, :roomuserbyofflinemutedmuted, :roomuserbybannedmutedmuted"
	cmds := strings.ReplaceAll(msg, ",", "\r")
	client.SendMOTDMessage(cmds)
}
