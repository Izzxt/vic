package core

type HabboRoomUnitStatus string

const (
	HabboRoomUnitStatusMove HabboRoomUnitStatus = "mv"

	HabboRoomUnitStatusSitIn  = "sit-in"
	HabboRoomUnitStatusSit    = "sit"
	HabboRoomUnitStatusSitOut = "sit-out"

	HabboRoomUnitStatusLayIn  = "lay-in"
	HabboRoomUnitStatusLay    = "lay"
	HabboRoomUnitStatusLayOut = "lay-out"

	HabboRoomUnitStatusFlatControl = "flatctrl"
	HabboRoomUnitStatusSign        = "sign"
	HabboRoomUnitStatusGesture     = "gst"
	HabboRoomUnitStatusWave        = "wav"
	HabboRoomUnitStatusTrading     = "trd"

	HabboRoomUnitStatusDip = "dip"

	HabboRoomUnitStatusEatIn  = "eat-in"
	HabboRoomUnitStatusEat    = "eat"
	HabboRoomUnitStatusEatOut = "eat-out"

	HabboRoomUnitStatusBeg = "beg"

	HabboRoomUnitStatusDeadIn  = "ded-in"
	HabboRoomUnitStatusDead    = "ded"
	HabboRoomUnitStatusDeadOut = "ded-out"

	HabboRoomUnitStatusJumpIn  = "jmp-in"
	HabboRoomUnitStatusJump    = "jmp"
	HabboRoomUnitStatusJumpOut = "jmp-out"

	HabboRoomUnitStatusPlayIn  = "pla-in"
	HabboRoomUnitStatusPlay    = "pla"
	HabboRoomUnitStatusPlayOut = "pla-out"

	HabboRoomUnitStatusSpeak   = "spk"
	HabboRoomUnitStatusCroak   = "crk"
	HabboRoomUnitStatusRelax   = "rlx"
	HabboRoomUnitStatusWings   = "wng"
	HabboRoomUnitStatusFlame   = "flm"
	HabboRoomUnitStatusRip     = "rip"
	HabboRoomUnitStatusGrow    = "grw"
	HabboRoomUnitStatusGrow1   = "grw1"
	HabboRoomUnitStatusGrow2   = "grw2"
	HabboRoomUnitStatusGrow3   = "grw3"
	HabboRoomUnitStatusGrow4   = "grw4"
	HabboRoomUnitStatusGrow5   = "grw5"
	HabboRoomUnitStatusGrow6   = "grw6"
	HabboRoomUnitStatusGrow7   = "grw7"
	HabboRoomUnitStatusKick    = "kck"
	HabboRoomUnitStatusWagTail = "wag"
	HabboRoomUnitStatusDance   = "dan"
	HabboRoomUnitStatusAms     = "ams"
	HabboRoomUnitStatusSwim    = "swm"
	HabboRoomUnitStatusTurn    = "trn"

	HabboRoomUnitStatusSrp   = "srp"
	HabboRoomUnitStatusSrpIn = "srp-in"

	HabboRoomUnitStatusSleepIn  = "slp-in"
	HabboRoomUnitStatusSleep    = "slp"
	HabboRoomUnitStatusSleepOut = "slp-out"
)

type NavigatorCategoryType string

const (
	NavigatorCategoryTypeOfficial = "official_view"
	NavigatorCategoryTypeHotel    = "hotel_view"
	NavigatorCategoryTypeRoomAds  = "roomads_view"
	NavigatorCategoryTypeMyWorld  = "myworld_view"
)

type NavigatorDisplayMode bool

const (
	NavigatorDisplayModeCollapsed NavigatorDisplayMode = true
	NavigatorDisplayModeExpanded                       = false
)

type NavigatorListMode int32

const (
	NavigatorListModeList NavigatorListMode = iota
	NavigatorListModeThumbnail
	NavigatorListModeForceThumbnail
)

type NavigatorDisplayOrder int32

const (
	NavigatorDisplayOrderOrder NavigatorDisplayOrder = iota
	NavigatorDisplayOrderActivity
)

type NavigatorSearchAction int32

const (
	NavigatorSearchActionNone NavigatorSearchAction = iota
	NavigatorSearchActionSearch
	NavigatorSearchActionGoToRoom
)
