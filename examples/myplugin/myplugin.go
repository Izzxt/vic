package main

import (
	"C"
	"fmt"

	"embed"
)
import "github.com/Izzxt/vic/core"

//go:embed manifest.json
var Manifest embed.FS

type myPlugin struct {
	client core.HabboClient
}

// OnLoad implements extensions.Plugin.
func (*myPlugin) OnLoad() {
	fmt.Println("OnLoad")
}

// OnStart implements extensions.Plugin.
func (*myPlugin) OnStart() {
	fmt.Println("OnStart")
}

// OnUnload implements extensions.Plugin.
func (*myPlugin) OnUnload() {
	fmt.Println("OnUnload")
}

// export New
func New(client core.HabboClient) core.Plugin {
	return &myPlugin{}
}
