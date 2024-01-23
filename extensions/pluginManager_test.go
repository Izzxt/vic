package extensions

import (
	"testing"
)

func TestWalkPlugins(t *testing.T) {
	plugin := NewPluginManager()
	plugin.LoadPlugin()
}
