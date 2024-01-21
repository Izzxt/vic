package bubbleschat

import (
	"context"
	"fmt"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/database"
	bubbles_chat "github.com/Izzxt/vic/database/rooms/bubbles_chat/querier"
)

type bubblesChat struct {
	ctx         context.Context
	bubblesChat bubbles_chat.BubblesChat
	bubbles     map[int32]bubbles_chat.BubblesChat
}

// GetBubbleChatById implements core.IBubblesChat.
func (b *bubblesChat) GetBubbleChatById(id int32) *bubbles_chat.BubblesChat {
	if bubble, ok := b.bubbles[id]; ok {
		return &bubble
	}
	return nil
}

// GetBubbleChatByRoomKey implements core.IBubblesChat.
func (b *bubblesChat) GetBubbleChatByKey(roomKey string) *bubbles_chat.BubblesChat {
	for _, bubble := range b.bubbles {
		if bubble.Key == roomKey {
			return &bubble
		}
	}
	return nil
}

func (b *bubblesChat) loadAll() {
	db := database.GetInstance().BubblesChat()
	bubbles, err := db.ListsBubbleChat(b.ctx)
	if err != nil {
		fmt.Printf("Error loading bubble chats: %v\n", err)
	}
	for _, bubble := range bubbles {
		b.bubbles[bubble.ID] = bubble
	}
}

func NewBubblesChat(ctx context.Context, bubbles bubbles_chat.BubblesChat) core.BubblesChat {
	bc := bubblesChat{ctx: ctx, bubblesChat: bubbles, bubbles: make(map[int32]bubbles_chat.BubblesChat)}
	bc.loadAll()
	return &bc
}
