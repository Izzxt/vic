package extensions

import (
	"fmt"
	"log"
	"path/filepath"
	"plugin"

	"github.com/Izzxt/vic/core"
)

type pluginManager struct{}

// LoadPlugin implements core.PluginManager.
func (p *pluginManager) LoadPlugin() {
	p.walkPlugins()

}

// walkPlugins implements core.PluginManager.
func (*pluginManager) walkPlugins() {
	matches, err := filepath.Glob("plugins/**/*.so")
	if err != nil {
		panic(err)
	}

	for _, match := range matches {
		fmt.Println(match)
		p, err := plugin.Open(match)
		if err != nil {
			log.Fatalf("Failed to open plugin: %v", err)
		}

		sym, err := p.Lookup("New")
		if err != nil {
			log.Fatalf("Failed to lookup symbol: %v", err)
		}

		newFunc, ok := sym.(func() Plugin)
		if !ok {
			log.Fatal("Plugin does not implement 'New' with interface 'extensions.Plugin' function")
		}

		plugin := newFunc()
		plugin.OnLoad()
		plugin.OnStart()
		plugin.OnUnload()
	}
}

// UnloadPlugin implements core.PluginManager.
func (*pluginManager) UnloadPlugin() {
	panic("unimplemented")
}

// RegisterPlugin implements core.PluginManager.
func (*pluginManager) RegisterPlugin() {
	panic("unimplemented")
}

func NewPluginManager() core.PluginManager {
	return &pluginManager{}
}
