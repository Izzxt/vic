// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package navigator_flat_cats

import ()

type NavigatorFlatCat struct {
	ID          int32  `json:"id"`
	MinRank     int32  `json:"min_rank"`
	CaptionSave string `json:"caption_save"`
	Caption     string `json:"caption"`
	AllowTrade  bool   `json:"allow_trade"`
	MaxUsers    int32  `json:"max_users"`
	IsPublic    bool   `json:"is_public"`
	OrderNum    int32  `json:"order_num"`
}
