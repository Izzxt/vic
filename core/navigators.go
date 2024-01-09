package core

import (
	navigator_flat_cats "github.com/Izzxt/vic/database/navigator/navigator_flat_cats/querier"
	navigator_public_cats "github.com/Izzxt/vic/database/navigator/navigator_public_cats/querier"
)

type INavigatorManager interface {
	NavigatorFlatCats() INavigatorFlatCats
	NavigatorPublicCats() INavigatorPublicCats
}

type INavigatorFlatCats interface {
	GetCategories() []navigator_flat_cats.NavigatorFlatCat
}

type INavigatorPublicCats interface {
	GetCategories() []navigator_public_cats.NavigatorPublicCat
}
