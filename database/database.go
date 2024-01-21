package database

import (
	"database/sql"
	"sync"

	navigator_flat_cats "github.com/Izzxt/vic/database/navigator/navigator_flat_cats/querier"
	navigator_public_cats "github.com/Izzxt/vic/database/navigator/navigator_public_cats/querier"
	bubbles_chat "github.com/Izzxt/vic/database/rooms/bubbles_chat/querier"
	room_info "github.com/Izzxt/vic/database/rooms/room_info/querier"
	room_models "github.com/Izzxt/vic/database/rooms/room_models/querier"
	users "github.com/Izzxt/vic/database/users/querier"
	users_stats "github.com/Izzxt/vic/database/users/stats/querier"
)

//go:generate sqlc generate -f ./database.json

var (
	instance *database
	once     sync.Once
)

type Queries interface {
	Users() users.Querier
	UsersStats() users_stats.Querier
	RoomInfo() room_info.Querier
	RoomModels() room_models.Querier
	NavigatorFlatCategories() navigator_flat_cats.Querier
	NavigatorPublicCategories() navigator_public_cats.Querier
	BubblesChat() bubbles_chat.Querier
	Close() error
}

type database struct {
	users                 users.Querier
	users_stats           users_stats.Querier
	room_info             room_info.Querier
	room_models           room_models.Querier
	navigator_flat_cats   navigator_flat_cats.Querier
	navigator_public_cats navigator_public_cats.Querier
	bubbles_chat          bubbles_chat.Querier
}

// UsersStats implements Queries.
func (db *database) UsersStats() users_stats.Querier {
	return db.users_stats
}

// NavigatorFlatCategories implements Queries.
func (db *database) NavigatorFlatCategories() navigator_flat_cats.Querier {
	return db.navigator_flat_cats
}

// NavigatorPublicCategories implements Queries.
func (db *database) NavigatorPublicCategories() navigator_public_cats.Querier {
	return db.navigator_public_cats
}

// RoomInfo implements Queries.
func (db *database) RoomInfo() room_info.Querier {
	return db.room_info
}

// RoomModels implements Queries.
func (db *database) RoomModels() room_models.Querier {
	return db.room_models
}

// Users implements Database.
func (db *database) Users() users.Querier {
	return db.users
}

// BubblesChat implements Database.
func (db *database) BubblesChat() bubbles_chat.Querier {
	return db.bubbles_chat
}

func GetInstance() Queries {
	return instance
}

func (db *database) Close() error {
	return db.Close()
}

func Init(db *sql.DB) Queries {
	if instance == nil {
		once.Do(func() {
			println("Initializing database")
			instance = &database{
				users:                 users.New(db),
				room_info:             room_info.New(db),
				room_models:           room_models.New(db),
				navigator_flat_cats:   navigator_flat_cats.New(db),
				navigator_public_cats: navigator_public_cats.New(db),
				users_stats:           users_stats.New(db),
				bubbles_chat:          bubbles_chat.New(db),
			}
		})
	}
	return instance
}
