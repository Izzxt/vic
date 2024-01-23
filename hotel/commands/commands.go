package commands

import (
	"strings"

	"github.com/Izzxt/vic/core"
	commands "github.com/Izzxt/vic/hotel/commands/common"
	room_chat "github.com/Izzxt/vic/packets/outgoing/room/units/chats"
)

type Commands struct {
	commands map[string]core.Command
}

func (c *Commands) Get(name string) core.Command {
	return c.commands[name]
}

func (c *Commands) Exists(name string) bool {
	_, ok := c.commands[name]
	return ok
}

// HandleCommand handles a command.
func (c *Commands) HandleCommand(client core.HabboClient, message string) {
	if strings.HasPrefix(message, ":") {
		message = message[1:]
		split := strings.Split(message, " ")
		commandName := split[0]
		args := split[1:]
		if !c.Exists(commandName) {
			client.Send(&room_chat.RoomUnitChatWhisperComposer{
				RoomUnit: client.GetHabbo().RoomUnit(), Message: "Command not found.",
				Gesture: 0, Bubble: 0, MessageLength: 0})
			return
		}
		command := c.Get(commandName)
		command.Execute(client, args)
	}
}

func (c *Commands) Register(name string, command core.Command) {
	c.commands[name] = command
}

// RegisterCommands registers all commands.
func (c *Commands) RegisterCommands() {
	c.Register("cmds", &commands.CmdCommand{})
	c.Register("about", &commands.AboutCommand{})
}

func NewCommandManager() core.CommandManager {
	return &Commands{commands: make(map[string]core.Command)}
}
