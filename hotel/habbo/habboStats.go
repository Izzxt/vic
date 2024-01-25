package habbo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/database"
	users_stats "github.com/Izzxt/vic/database/users/stats/querier"
)

type habboStats struct {
	ctx context.Context
	users_stats.UsersStat
}

// GetBubbleChat implements core.HabboStats.
func (h *habboStats) GetBubbleChat() int32 {
	return h.BubbleChatID.Int32
}

func (h *habboStats) Context() context.Context {
	return h.ctx
}

func (h *habboStats) Load(userId int32) core.HabboStats {
	db := database.GetInstance().UsersStats()
	stats, err := db.GetUserStats(h.ctx, userId)
	if err != nil {
		_, err := db.InsertUserStats(h.ctx, userId)
		if err != nil {
			fmt.Printf("Error inserting user stats: %v\n", err)
		}
		stats, err := db.GetUserStats(h.ctx, userId)
		if err != nil {
			fmt.Printf("Error loading user stats: %v\n", err)
		}
		h.UsersStat = stats
	}
	h.UsersStat = stats
	return h
}

// UpdateBubbleChat implements core.IHabboStats.
func (h *habboStats) UpdateBubbleChat(styleId int32) {
	db := database.GetInstance().UsersStats()
	if err := db.UpdateBubbleChat(h.ctx, users_stats.UpdateBubbleChatParams{
		UserID:       h.UserID,
		BubbleChatID: sql.NullInt32{Int32: styleId, Valid: true},
	}); err != nil {
		fmt.Printf("Error updating bubble chat: %v\n", err)
	}

	h.BubbleChatID = sql.NullInt32{Int32: styleId, Valid: true}
}

func NewHabboStats(ctx context.Context, user users_stats.UsersStat) core.HabboStats {
	return &habboStats{ctx: ctx, UsersStat: user}
}
