package incoming

const (
	ReleaseVersionEvent = 4000
	SecureLoginEvent    = 2419 // 1989
	UniqueIdEvent       = 2490 // 3521
	VersionCheckEvent   = 1053 // 1220
	PingEvent           = 295  // 878
	PongEvent           = 2596

	// users
	RequestUserDataEvent       = 357 // 2629
	RequestUserCreditsEvent    = 273 // 1051
	RequestUserClubEvent       = 3166
	RequestMeMenuSettingsEvent = 2388
	UsernameEvent              = 3878
	UserFigure                 = 2730

	// Navigator
	NewNavigatorEvent             = 2110 // 3375 navigator
	RequestRoomCategoriesEvent    = 3027
	RequestNewNavigatorRoomsEvent = 249

	EventTrackerEvent = 3457 // 143

	// Hotel view
	RequestHotelViewBonusRareEvent = 957
	HotelViewDataEvent             = 2912

	// Friends
	RequestFriendsEvent     = 1523
	RequestInitFriendsEvent = 2781

	// Room
	RequestRoomDataEvent      = 2230
	RequestRoomLoadEvent      = 2312
	RequestRoomHeightmapEvent = 3898
	RequestHeightmapEvent     = 2300
	RoomCreateEvent           = 2752

	// Room unit
	RoomUnitWalkEvent        = 3320
	RoomUserStartTypingEvent = 1597
	RoomUserStopTypingEvent  = 1474
	RoomUnitChatEvent        = 1314
	RoomUnitChatShoutEvent   = 2085
	RoomUnitChatWhisperEvent = 1543
)
