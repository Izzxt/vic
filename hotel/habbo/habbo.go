package habbo

import "github.com/Izzxt/vic/core"

type habbo struct{}

func NewHabbo() core.IHabbo {
	return &habbo{}
}
