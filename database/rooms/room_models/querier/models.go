// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package room_models

import ()

type RoomModel struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	Heightmap string `json:"heightmap"`
	IsClub    bool   `json:"is_club"`
	IsCustom  bool   `json:"is_custom"`
	X         int32  `json:"x"`
	Y         int32  `json:"y"`
	Dir       int32  `json:"dir"`
}
