package extensions

import (
	"encoding/json"
	"os"
)

type Plugin interface {
	OnStart()
	OnLoad()
	OnUnload()
}

type pluginConfig struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

// GetAuthor implements PluginConfig.
func (p *pluginConfig) GetAuthor() string {
	return p.Author
}

// GetDescription implements PluginConfig.
func (p *pluginConfig) GetDescription() string {
	return p.Description
}

// GetName implements PluginConfig.
func (p *pluginConfig) GetName() string {
	return p.Name
}

// GetVersion implements PluginConfig.
func (p *pluginConfig) GetVersion() string {
	return p.Version
}

type PluginConfig interface {
	GetName() string
	GetVersion() string
	GetAuthor() string
	GetDescription() string
	LoadConfigFile(string) (PluginConfig, error)
}

// LoadConfigFile implements PluginConfig.
func (p *pluginConfig) LoadConfigFile(path string) (PluginConfig, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, p); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *pluginConfig) UnmarshalJSON(data []byte) error {
	type Alias pluginConfig
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

func NewPluginConfig() PluginConfig {
	return &pluginConfig{}
}
