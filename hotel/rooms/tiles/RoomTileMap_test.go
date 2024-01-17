package tiles_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Izzxt/vic/core"
	"github.com/Izzxt/vic/hotel/rooms"
	"github.com/Izzxt/vic/hotel/rooms/tiles"
)

func TestRoomTileMap_PathFinder(t *testing.T) {
	type args struct {
		start core.RoomTile
		goal  core.RoomTile
	}
	tests := []struct {
		name string
		args args
		want []core.RoomTile
	}{
		{
			name: "test",
			args: args{
				start: tiles.NewRoomTile(0, 0, 33, 2),
				goal:  tiles.NewRoomTile(5, 7, 0, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			roomModels := &rooms.RoomModels{}
			roomModels.Heightmap = "xxxxxx\n\rx00000\n\rx00000\n\r000000\n\rx00000\n\rx00000\n\rx00000\n\rx00000"
			tileMap := tiles.NewRoomTileMap(nil, roomModels)
			for _, row := range tileMap.GetTiles() {
				for _, tile := range row {
					fmt.Printf("| X: %d Y: %d H: %d S: %d \t", tile.GetX(), tile.GetY(), tile.GetHeight(), tile.GetState())
				}
				fmt.Println()
			}
			if got := tileMap.FindPath(tt.args.start, tt.args.goal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoomTileMap.PathFinder() = %v, want %v", got, tt.want)
			}
		})
	}
}
