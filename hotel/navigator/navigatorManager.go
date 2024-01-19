package navigator

import (
	"context"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/packets/outgoing/navigator"
)

type navigatorManager struct {
	navigatorFlatCats   core.NavigatorFlatCats
	navigatorPublicCats core.NavigatorPublicCats
	searchResultLists   map[string]core.NavigatorSearchResults
}

func (nm *navigatorManager) NavigatorFlatCats() core.NavigatorFlatCats {
	return nm.navigatorFlatCats
}

func (nm navigatorManager) NavigatorPublicCats() core.NavigatorPublicCats {
	return nm.navigatorPublicCats
}

func (nm navigatorManager) SearchCategory(client core.HabboClient, categoryType string) []core.NavigatorSearchResults {
	room := client.Room()
	results := make([]core.NavigatorSearchResults, 0)

	switch categoryType {
	case core.NavigatorCategoryTypeMyWorld:
		rm := make([]interface{}, 0)
		rooms := room.GetRoomsByOwnerId(client.GetHabbo().HabboInfo().ID)
		for _, r := range rooms {
			rm = append(rm, r)
		}
		results = append(results, core.NavigatorSearchResults{
			Identifier: core.NavigatorCategoryTypeMyWorld,
			Rooms:      rm,
		})

		client.Send(
			&navigator.NewNavigatorSearchResultsComposer{
				SearchCode:             categoryType,
				SearchQuery:            "",
				NavigatorSearchResults: results,
			},
		)

	case core.NavigatorCategoryTypeHotel:
		rm := make([]interface{}, 0)
		rooms := room.GetActiveRooms()
		for _, r := range rooms {
			rm = append(rm, r)
		}
		results = append(results, core.NavigatorSearchResults{
			Identifier: core.NavigatorCategoryTypeHotel,
			Rooms:      rm,
		})

	case core.NavigatorCategoryTypeOfficial:
		// TODO: implement
	case core.NavigatorCategoryTypeRoomAds:
		// TODO: implement

	}

	return results
}

func NewNavigatorManager(ctx context.Context) core.NavigatorManager {
	navigatorFlatCats := NewNavigatorFlatCats(ctx)
	navigatorPublicCats := NewNavigatorPublicCats(ctx)

	return &navigatorManager{
		navigatorFlatCats:   navigatorFlatCats,
		navigatorPublicCats: navigatorPublicCats,
	}
}
