// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package navigator_flat_cats

import (
	"context"
)

type Querier interface {
	ListNavigatorFlatCategories(ctx context.Context) ([]NavigatorFlatCat, error)
}

var _ Querier = (*Queries)(nil)