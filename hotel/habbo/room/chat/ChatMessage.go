package chat_message

import (
	"strings"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/hotel/commands"
	"github.com/Izzxt/vic/packets/incoming"
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
	if strings.HasPrefix(c.message, ":") {
		c.isCommand = true
		c.command.HandleCommand(c.client, c.message)
	}
}

func (c *chatMessage) parseWhisperMessage() {
	target := strings.Split(c.message, " ")
	c.target = c.client.GetHabbo().Room().GetHabboByName(target[0])
	c.message = strings.Join(target[1:], " ")
}

func (c *chatMessage) SendMessage() {
	// TODO: check if muted
	// TODO: add gesture functionallity
	if c.isWhisper {
		c.parseWhisperMessage()

		if c.target == nil || c.client.GetHabbo() == c.target {
			c.client.GetHabbo().Whisper(c.client.GetHabbo(), c.message, 0)
			return
		}

		c.client.GetHabbo().Whisper(c.client.GetHabbo(), c.message, 0)
		c.target.Client().GetHabbo().Whisper(c.client.GetHabbo(), c.message, 0)
	} else if c.isShout {
		c.client.GetHabbo().Shout(c.message, 0)
	} else {
		c.parseCommand()
		if !c.isCommand {
			c.client.GetHabbo().Talk(c.message, c.styleId)
		}
	}
}

func NewChatMessage(client core.HabboClient, header int16, message string, styleId int32) core.ChatMessage {
	command := commands.NewCommandManager()
	command.RegisterCommands()

	if styleId != 1 {
		client.GetHabbo().HabboStats().UpdateBubbleChat(styleId)
	} else {
		styleId = client.GetHabbo().HabboStats().GetBubbleChat()
	}

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
