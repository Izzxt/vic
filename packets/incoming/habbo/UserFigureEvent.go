package habbo

import (
	"fmt"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/habbo"
)

type UserFigureEvent struct{}

func (u UserFigureEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {
	fmt.Printf("\033[28mHabbo: %v\n", client.GetHabbo())
	if client.GetHabbo() != nil {
		return
	}

	gender := in.ReadString()
	figure := in.ReadString()

	fmt.Println(gender, figure)

	client.Send(&habbo.UserFigureComposer{Habbo: client.GetHabbo()})
}
