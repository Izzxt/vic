package hotelview

import (
	"strings"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/hotelview"
)

type HotelViewDataEvent struct{}

// Execute implements core.IIncomingMessage.
func (*HotelViewDataEvent) Execute(client core.IHabboClient, in core.IIncomingPacket) {

	data := in.ReadString()
	if strings.Contains(data, ";") {
		datas := strings.Split(data, ";")

		for _, d := range datas {
			if strings.Contains(d, ",") {
				client.Send(&hotelview.HotelViewDataComposer{Data: d, Key: strings.Split(d, ",")[len(strings.Split(d, ","))-1]})
			} else {
				client.Send(&hotelview.HotelViewDataComposer{Data: data, Key: d})
			}
			break
		}
	} else {
		client.Send(&hotelview.HotelViewDataComposer{Data: data, Key: strings.Split(data, ",")[len(strings.Split(data, ","))-1]})

	}

}
