package navigator

import (
	"context"

	"github.com/Izzxt/vic/core"
)

type navigatorManager struct {
	navigatorFlatCats   core.NavigatorFlatCats
	navigatorPublicCats core.NavigatorPublicCats
}

func (nm *navigatorManager) NavigatorFlatCats() core.NavigatorFlatCats {
	return nm.navigatorFlatCats
}

func (nm navigatorManager) NavigatorPublicCats() core.NavigatorPublicCats {
	return nm.navigatorPublicCats
}

func NewNavigatorManager(ctx context.Context) core.NavigatorManager {
	navigatorFlatCats := NewNavigatorFlatCats(ctx)
	navigatorPublicCats := NewNavigatorPublicCats(ctx)
	return &navigatorManager{
		navigatorFlatCats:   navigatorFlatCats,
		navigatorPublicCats: navigatorPublicCats,
	}
}
