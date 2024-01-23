package chat_message

import (
	"strings"
	"unicode/utf8"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/hotel/commands"
	"github.com/Izzxt/vic/packets/incoming"
	room_chat "github.com/Izzxt/vic/packets/outgoing/room/units/chats"
)

type chatMessage struct {
	isWhisper bool
	isShout   bool
	target    core.Habbo
	client    core.HabboClient
	isCommand bool
	message   string
	styleId   int32
	command   core.CommandManager
}

func (c *chatMessage) parseCommand() {
	c.isCommand = true
	c.command.HandleCommand(c.client, c.message)
}

func (c *chatMessage) parseWhisperMessage() {
	target := strings.Split(c.message, " ")
	c.target = c.client.GetHabbo().Room().GetHabboByName(target[0])
	c.message = strings.Join(target[1:], " ")
}

func (c *chatMessage) SendMessage() {
	// TODO: check if muted
	// TODO: add bubble functionallity
	// TODO: add gesture functionallity
	if c.isWhisper {
		c.parseWhisperMessage()
		c.client.Send(&room_chat.RoomUnitChatWhisperComposer{
			RoomUnit: c.client.GetHabbo().RoomUnit(), Message: c.message, Gesture: 0,
			Bubble: 0, MessageLength: int32(utf8.RuneCountInString(c.message))})

		if c.target != nil {
			c.target.Client().Send(&room_chat.RoomUnitChatWhisperComposer{
				RoomUnit: c.client.GetHabbo().RoomUnit(), Message: c.message, Gesture: 0,
				Bubble: 0, MessageLength: int32(utf8.RuneCountInString(c.message))})
		}
	} else if c.isShout {
		c.client.SendToRoom(c.client.GetHabbo().Room(), &room_chat.RoomUnitChatShoutComposer{
			RoomUnit: c.client.GetHabbo().RoomUnit(), Message: c.message, Gesture: 0,
			Bubble: 0, MessageLength: int32(utf8.RuneCountInString(c.message))})
	} else {
		c.parseCommand()
		if c.isCommand {
			return
		}

		c.client.SendToRoom(c.client.GetHabbo().Room(), &room_chat.RoomUnitChatComposer{
			RoomUnit: c.client.GetHabbo().RoomUnit(), Message: c.message, Gesture: 0,
			Bubble: 0, MessageLength: int32(utf8.RuneCountInString(c.message))})
	}
}

func NewChatMessage(client core.HabboClient, header int16, message string, styleId int32) core.ChatMessage {
	command := commands.NewCommandManager()
	command.RegisterCommands()
	chatMsg := chatMessage{client: client, isWhisper: false, isShout: false, message: message,
		styleId: styleId, isCommand: false, command: command,
	}
	if header == incoming.RoomUnitChatShoutEvent {
		chatMsg.isShout = true
		return &chatMsg
	} else if header == incoming.RoomUnitChatWhisperEvent {
		chatMsg.isWhisper = true
		return &chatMsg
	}
	return &chatMsg
}
