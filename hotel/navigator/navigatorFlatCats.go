package navigator

import (
	"context"
	"log"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/database"
	navigator_flat_cats "github.com/Izzxt/vic/database/navigator/navigator_flat_cats/querier"
)

type navigatorFlatCats struct {
	ctx context.Context
	navigator_flat_cats.NavigatorFlatCat
}

func (n *navigatorFlatCats) GetCategories() []navigator_flat_cats.NavigatorFlatCat {
	db := database.GetInstance().NavigatorFlatCategories()

	cats, err := db.ListNavigatorFlatCategories(n.ctx)
	if err != nil {
		log.Fatalf("failed to get navigator flat categories: %v", err)
	}

	return cats
}

func NewNavigatorFlatCats(ctx context.Context) core.NavigatorFlatCats {
	return &navigatorFlatCats{ctx: ctx}
}
