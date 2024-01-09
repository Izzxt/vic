package navigator

import (
	"context"
	"log"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/database"
	navigator_public_cats "github.com/Izzxt/vic/database/navigator/navigator_public_cats/querier"
)

type navigatorPublicCats struct {
	ctx context.Context
	navigator_public_cats.NavigatorPublicCat
}

func (n *navigatorPublicCats) GetCategories() []navigator_public_cats.NavigatorPublicCat {
	db := database.GetInstance().NavigatorPublicCategories()

	cats, err := db.ListNavigatorPublicCategories(n.ctx)
	if err != nil {
		log.Fatalf("failed to get navigator public categories: %v", err)
	}

	return cats
}

func NewNavigatorPublicCats(ctx context.Context) core.INavigatorPublicCats {
	return &navigatorPublicCats{ctx: ctx}
}
