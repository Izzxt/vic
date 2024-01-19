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

	categories map[int32]navigator_flat_cats.NavigatorFlatCat
}

func (n *navigatorFlatCats) loadCategories() {
	db := database.GetInstance().NavigatorFlatCategories()
	cats, err := db.ListNavigatorFlatCategories(n.ctx)
	if err != nil {
		log.Fatalf("failed to get navigator flat categories: %v", err)
	}
	for _, cat := range cats {
		n.categories[cat.ID] = cat
	}
}

func (n *navigatorFlatCats) GetCategories() []navigator_flat_cats.NavigatorFlatCat {
	categories := make([]navigator_flat_cats.NavigatorFlatCat, 0)
	for _, cat := range n.categories {
		categories = append(categories, cat)
	}
	return categories
}

func (n *navigatorFlatCats) GetCategory(category int32) navigator_flat_cats.NavigatorFlatCat {
	var cat navigator_flat_cats.NavigatorFlatCat
	for _, c := range n.categories {
		if c.ID == category {
			cat = c
		}
	}
	return cat
}

func NewNavigatorFlatCats(ctx context.Context) core.NavigatorFlatCats {
	cats := &navigatorFlatCats{ctx: ctx}
	cats.categories = make(map[int32]navigator_flat_cats.NavigatorFlatCat)
	cats.loadCategories()
	return cats
}
