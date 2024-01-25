package extensions

import (
	"encoding/json"

	"github.com/Izzxt/vic/core"
)

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

// LoadConfigFile implements PluginConfig.
func (p *pluginConfig) LoadConfigFile(bytes []byte) (core.PluginConfig, error) {

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

func NewPluginConfig() core.PluginConfig {
	return &pluginConfig{}
}
