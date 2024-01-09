package navigator

import (
	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing"
)

type NavigatorSearchResults struct {
	Identifier string
	PublicName string
}

type NewNavigatorSearchResultsComposer struct {
	SearchCode             string
	SearchQuery            string
	NavigatorSearchResults []NavigatorSearchResults
}

// Compose implements core.IOutgoingMessage.
func (c *NewNavigatorSearchResultsComposer) Compose(compose core.IOutgoingPacket) core.IOutgoingPacket {
	compose.WriteString(c.SearchCode)
	compose.WriteString(c.SearchQuery)
	compose.WriteInt(int32(len(c.NavigatorSearchResults)))

	for range c.NavigatorSearchResults {
		compose.WriteString(c.SearchCode)
		compose.WriteString(c.SearchQuery)
		compose.WriteInt(0)
		compose.WriteBool(false)
		compose.WriteInt(0)
		compose.WriteInt(1)
		room(compose)
	}
	return compose
}

// TODO: move to room.go
func room(compose core.IOutgoingPacket) {
	var isPublic bool = false
	var tags []string
	compose.WriteInt(1)         // room id
	compose.WriteString("test") // room name

	if isPublic {
		compose.WriteInt(0)
		compose.WriteString("")
	} else {
		compose.WriteInt(1)          // room owner id
		compose.WriteString("Izzxt") // room owner name
	}

	compose.WriteInt(0) // door mode

	compose.WriteInt(0)  // user count
	compose.WriteInt(10) // max user count

	compose.WriteString("a room") // description

	compose.WriteInt(0) // trade mode
	compose.WriteInt(0) // score

	compose.WriteInt(0)
	compose.WriteInt(9) // category id

	compose.WriteInt(int32(len(tags))) // tag count
	for _, tag := range tags {
		compose.WriteString(tag)
	}

	var base = 0

	if !isPublic {
		base = base | 8
	}

	compose.WriteInt(int32(base))
}

// GetId implements core.IOutgoingMessage.
func (*NewNavigatorSearchResultsComposer) GetId() uint16 {
	return outgoing.NewNavigatorSearchResultsComposer
}
