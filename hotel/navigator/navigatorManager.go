package navigator

import (
	"context"

	"github.com/Izzxt/vic/core"
)

type navigatorManager struct {
	navigatorFlatCats   core.INavigatorFlatCats
	navigatorPublicCats core.INavigatorPublicCats
}

func (nm *navigatorManager) NavigatorFlatCats() core.INavigatorFlatCats {
	return nm.navigatorFlatCats
}

func (nm navigatorManager) NavigatorPublicCats() core.INavigatorPublicCats {
	return nm.navigatorPublicCats
}

func NewNavigatorManager(ctx context.Context) core.INavigatorManager {
	navigatorFlatCats := NewNavigatorFlatCats(ctx)
	navigatorPublicCats := NewNavigatorPublicCats(ctx)
	return &navigatorManager{
		navigatorFlatCats:   navigatorFlatCats,
		navigatorPublicCats: navigatorPublicCats,
	}
}
