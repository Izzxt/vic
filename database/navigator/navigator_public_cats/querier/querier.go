// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package navigator_public_cats

import (
	"context"
)

type Querier interface {
	ListNavigatorPublicCategories(ctx context.Context) ([]NavigatorPublicCat, error)
}

var _ Querier = (*Queries)(nil)
