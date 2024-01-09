package rooms

import (
	"context"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/database"
	room_models "github.com/Izzxt/vic/database/rooms/room_models/querier"
)

type RoomModels struct {
	room_models.RoomModel
	ctx context.Context
}

// GetDir implements core.IRoomModel.
func (r *RoomModels) GetDir() int32 {
	return r.RoomModel.Dir
}

// GetHeightmap implements core.IRoomModel.
func (r *RoomModels) GetHeightmap() string {
	return r.RoomModel.Heightmap
}

// GetId implements core.IRoomModel.
func (r *RoomModels) GetId() int32 {
	return r.RoomModel.ID
}

// GetIsClub implements core.IRoomModel.
func (r *RoomModels) GetIsClub() bool {
	return r.RoomModel.IsClub
}

// GetIsCustom implements core.IRoomModel.
func (r *RoomModels) GetIsCustom() bool {
	return r.RoomModel.IsCustom
}

// GetName implements core.IRoomModel.
func (r *RoomModels) GetName() string {
	return r.RoomModel.Name
}

// GetX implements core.IRoomModel.
func (r *RoomModels) GetX() int32 {
	return r.RoomModel.X
}

// GetY implements core.IRoomModel.
func (r *RoomModels) GetY() int32 {
	return r.RoomModel.Y
}

func (r *RoomModels) Load(id int32) core.IRoomModel {
	db := database.GetInstance().RoomModels()

	model, err := db.GetModelById(r.ctx, id)
	if err != nil {
		panic(err)
	}

	r.RoomModel = model
	return r
}

func NewRoomModels(ctx context.Context) core.IRoomModel {
	return &RoomModels{ctx: ctx}
}
