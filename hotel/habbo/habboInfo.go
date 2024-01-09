package habbo

import (
	"context"

	users "github.com/Izzxt/vic/database/users/querier"
)

type habboInfo struct {
	ctx context.Context
	users.User
}

func NewHabboInfo(ctx context.Context, user users.User) *habboInfo {
	return &habboInfo{ctx: ctx, User: user}
}
