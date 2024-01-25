package extensions

import (
	"embed"
	"log"
	"path/filepath"
	"plugin"

	"github.com/Izzxt/vic/core"
)

type pluginManager struct {
	config  core.PluginConfig
	plugins map[core.PluginConfig]core.Plugin
	client  core.HabboClient
}

func (m *pluginManager) loadPlugin() {
	pluginMatches, err := filepath.Glob("plugins/*.so")
	if err != nil {
		panic(err)
	}

	for _, match := range pluginMatches {
		p, err := plugin.Open(match)
		if err != nil {
			log.Fatalf("Failed to open plugin: %v", err)
		}

		manifestSym, err := p.Lookup("Manifest")
		if err != nil {
			log.Fatalf("Failed to lookup symbol: %v", err)
		}

		manifest, ok := manifestSym.(*embed.FS)
		if !ok {
			log.Fatal("Plugin does not have 'Manifest' variables with type 'embed.FS'")
		}

		manifestByte, err := manifest.ReadFile("manifest.json")
		if err != nil {
			log.Fatalf("Failed to open manifest file: %v", err)
		}

		config, err := m.config.LoadConfigFile(manifestByte)
		if err != nil {
			log.Fatalf("Failed to load config file: %v", err)
		}

		sym, err := p.Lookup("New")
		if err != nil {
			log.Fatalf("Failed to lookup symbol: %v", err)
		}

		newFunc, ok := sym.(func(core.HabboClient) core.Plugin)
		if !ok {
			log.Fatal("Plugin does not implement 'New' with interface 'extensions.Plugin' function")
		}

		plugin := newFunc(m.client)

		m.plugins[config] = plugin
	}
}

// LoadPlugin implements core.PluginManager.
func (m *pluginManager) LoadPlugin() {
	for _, plugin := range m.plugins {
		plugin.OnLoad()
	}
}

func (m *pluginManager) StartPlugin() {
	for _, plugin := range m.plugins {
		plugin.OnStart()
	}
}

// UnloadPlugin implements core.PluginManager.
func (m *pluginManager) UnloadPlugin() {
	for _, plugin := range m.plugins {
		plugin.OnUnload()
	}
}

// GetPluginByName implements core.PluginManager.
func (m *pluginManager) GetPluginByName(name string) core.Plugin {
	for config, plugin := range m.plugins {
		if config.GetName() == name {
			return plugin
		}
	}
	return nil
}

// SetClient implements core.PluginManager.
func (m *pluginManager) SetClient(client core.HabboClient) {
	m.client = client
}

func NewPluginManager() core.PluginManager {
	m := &pluginManager{}
	m.config = NewPluginConfig()
	m.plugins = make(map[core.PluginConfig]core.Plugin)
	m.loadPlugin()
	return m
}
