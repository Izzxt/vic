package codec

import (
	"bytes"
	"log"

	"github.com/Izzxt/vic/core"
	"github.com/gorilla/websocket"
)

var XmlPolicy = []byte("<?xml version=\"1.0\"?>\r\n" +
	"<!DOCTYPE cross-domain-policy SYSTEM \"/xml/dtds/cross-domain-policy.dtd\">\r\n" +
	"<cross-domain-policy>\r\n" +
	"<allow-access-from domain=\"*\" to-ports=\"1-31111\" />\r\n" +
	"</cross-domain-policy>\x00")

func SendPolicy(buf bytes.Buffer, client core.HabboClient) {
	// var hasReceivedPolicy = false
	b, err := buf.ReadByte()
	if err != nil {
		log.Fatalf("failed to read byte: %v", err)
	}
	if b == '<' {
		// hasReceivedPolicy = true
		err := client.Connection().WriteMessage(websocket.BinaryMessage, XmlPolicy)
		if err != nil {
			log.Fatalf("failed to write xml policy: %v", err)
		}
	}
}
