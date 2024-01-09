package database

import (
	"database/sql"
	"sync"

	navigator_flat_cats "github.com/Izzxt/vic/database/navigator/navigator_flat_cats/querier"
	navigator_public_cats "github.com/Izzxt/vic/database/navigator/navigator_public_cats/querier"
	room_info "github.com/Izzxt/vic/database/rooms/room_info/querier"
	room_models "github.com/Izzxt/vic/database/rooms/room_models/querier"
	users "github.com/Izzxt/vic/database/users/querier"
)

//go:generate sqlc generate -f ./database.json

var (
	instance *database
	once     sync.Once
)

type Queries interface {
	Users() users.Querier
	RoomInfo() room_info.Querier
	RoomModels() room_models.Querier
	NavigatorFlatCategories() navigator_flat_cats.Querier
	NavigatorPublicCategories() navigator_public_cats.Querier
}

type database struct {
	users                 users.Querier
	room_info             room_info.Querier
	room_models           room_models.Querier
	navigator_flat_cats   navigator_flat_cats.Querier
	navigator_public_cats navigator_public_cats.Querier
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

func GetInstance() Queries {
	return instance
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
			}
		})
	}
	return instance
}
