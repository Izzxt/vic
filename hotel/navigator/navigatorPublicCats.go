package navigator

import (
	"context"
	"fmt"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/database"
	navigator_public_cats "github.com/Izzxt/vic/database/navigator/navigator_public_cats/querier"
)

type navigatorPublicCats struct {
	ctx context.Context
	navigator_public_cats.NavigatorPublicCat

	categories map[int32]navigator_public_cats.NavigatorPublicCat
}

func (n *navigatorPublicCats) loadCategories() {
	db := database.GetInstance().NavigatorPublicCategories()
	cats, err := db.ListNavigatorPublicCategories(n.ctx)
	if err != nil {
		fmt.Printf("failed to get navigator public categories: %v", err)
	}
	for _, cat := range cats {
		n.categories[cat.ID] = cat
	}
}

func (n *navigatorPublicCats) GetCategories() []navigator_public_cats.NavigatorPublicCat {
	categories := make([]navigator_public_cats.NavigatorPublicCat, 0)
	for _, cat := range n.categories {
		categories = append(categories, cat)
	}
	return categories
}

func (n *navigatorPublicCats) GetCategory(category int32) navigator_public_cats.NavigatorPublicCat {
	var cat navigator_public_cats.NavigatorPublicCat
	for _, c := range n.categories {
		if c.ID == category {
			cat = c
		}
	}
	return cat
}

func NewNavigatorPublicCats(ctx context.Context) core.NavigatorPublicCats {
	cats := &navigatorPublicCats{ctx: ctx}
	cats.categories = make(map[int32]navigator_public_cats.NavigatorPublicCat)
	cats.loadCategories()
	return cats
}
