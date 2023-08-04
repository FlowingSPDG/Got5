package models

// https://splewis.github.io/get5/latest/events_and_forwards/#sourcemod-forwards

// All Events

// OnEventPayload Called when any event is fired. This forward takes two parameters (while all others take only one); the event object itself and a string representing the encoded JSON object. If you use this forward, you should fetch the event name and switch on it, casting the object to its correct subclass using
type OnEventPayload map[string]any

type Event struct {
	Event string `json:"event"`
}

// Series Flow

// OnGameStateChangedPayload Events the occur in relation to setting up a match or series.
type OnGameStateChangedPayload struct {
	Event
	NewState string `json:"new_state"`
	OldState string `json:"old_state"`
}

// OnPreLoadMatchConfigPayload Fired when the server attempts to load a match config.
type OnPreLoadMatchConfigPayload struct {
	Event
	Filename string `json:"filename"`
}

// OnLoadMatchConfigFailedPayload Fired when a match config fails to load.
type OnLoadMatchConfigFailedPayload struct {
	Event
	Reason string `json:"reason"`
}

// OnSeriesInitPayload Fired when a series is started after loading a match config.
type OnSeriesInitPayload struct {
	Event
	MatchID string `json:"matchid"`
	NumMaps int    `json:"num_maps"`
	Team1   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"team1"`
	Team2 struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"team2"`
}

// OnMapResultPayload Fired when the map ends.
type OnMapResultPayload struct {
	Event
	MatchID   string        `json:"matchid"`
	MapNumber int           `json:"map_number"`
	Team1     Get5StatsTeam `json:"team1"`
	Team2     Get5StatsTeam `json:"team2"`
	Winner    Winner        `json:"winner"`
}

type Get5StatsTeam struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	SeriesScore  int               `json:"series_score"`
	Score        int               `json:"score"`
	ScoreCt      int               `json:"score_ct"`
	ScoreT       int               `json:"score_t"`
	Players      []Get5StatsPlayer `json:"players"`
	Side         string            `json:"side"`
	StartingSide string            `json:"starting_side"`
}

type Get5StatsPlayer struct {
	SteamID string `json:"steamid"`
	Name    string `json:"name"`
	Stats   struct {
		Kills             int `json:"kills"`
		Deaths            int `json:"deaths"`
		Assists           int `json:"assists"`
		FlashAssists      int `json:"flash_assists"`
		TeamKills         int `json:"team_kills"`
		Suicides          int `json:"suicides"`
		Damage            int `json:"damage"`
		UtilityDamage     int `json:"utility_damage"`
		EnemiesFlashed    int `json:"enemies_flashed"`
		FriendliesFlashed int `json:"friendlies_flashed"`
		KnifeKills        int `json:"knife_kills"`
		HeadshotKills     int `json:"headshot_kills"`
		RoundsPlayed      int `json:"rounds_played"`
		BombDefuses       int `json:"bomb_defuses"`
		BombPlants        int `json:"bomb_plants"`
		OneK              int `json:"1k"`
		TwoK              int `json:"2k"`
		ThreeK            int `json:"3k"`
		FourK             int `json:"4k"`
		FiveK             int `json:"5k"`
		OneV1             int `json:"1v1"`
		OneV2             int `json:"1v2"`
		OneV3             int `json:"1v3"`
		OneV4             int `json:"1v4"`
		OneV5             int `json:"1v5"`
		FirstKillsT       int `json:"first_kills_t"`
		FirstKillsCt      int `json:"first_kills_ct"`
		FirstDeathsT      int `json:"first_deaths_t"`
		FirstDeathsCt     int `json:"first_deaths_ct"`
		TradeKills        int `json:"trade_kills"`
		KAST              int `json:"kast"`
		Score             int `json:"score"`
		Mvp               int `json:"mvp"`
	} `json:"stats"`
}

// OnSeriesResultPayload Fired when a series is over. winner indicates team and side 0 if there was no winner in cases of a draw or if the series was forcefully canceled.
type OnSeriesResultPayload struct {
	Event
	MatchID          string `json:"matchid"`
	Team1SeriesScore int    `json:"team1_series_score"`
	Team2SeriesScore int    `json:"team2_series_score"`
	Winner           Winner `json:"winner"`
	TimeUntilRestore int    `json:"time_until_restore"`
}

// OnSidePickedPayload Fired when a side is picked by a team.
type OnSidePickedPayload struct {
	Event
	MatchID   string `json:"matchid"`
	Team      string `json:"team"`
	MapName   string `json:"map_name"`
	Side      string `json:"side"`
	MapNumber int    `json:"map_number"`
}

// OnMapPickedPayload Fired when a team picks a map.
type OnMapPickedPayload struct {
	Event
	MatchID   string `json:"matchid"`
	Team      string `json:"team"`
	MapName   string `json:"map_name"`
	MapNumber int    `json:"map_number"`
}

// OnMapVetoedPayload Fired when a team vetos a map.
type OnMapVetoedPayload struct {
	Event
	MatchID string `json:"matchid"`
	Team    string `json:"team"`
	MapName string `json:"map_name"`
}

// OnBackupRestorePayload Fired when a round is restored from a backup. Note that the map and round numbers indicate the round being restored to, not the round the backup was requested during.
type OnBackupRestorePayload struct {
	Event
	MatchID     string `json:"matchid"`
	MapNumber   int    `json:"map_number"`
	RoundNumber int    `json:"round_number"`
	Filename    string `json:"filename"`
}

// OnDemoFinishedPayload Fired when the GOTV recording has ended. This event does not fire if no demo was recorded.
type OnDemoFinishedPayload struct {
	Event
	MatchID   string `json:"matchid"`
	MapNumber int    `json:"map_number"`
	Filename  string `json:"filename"`
}

// OnDemoUploadEndedPayload Fired when the request to upload a demo ends, regardless if it succeeds or fails. If you upload demos, you should not shut down a server until this event has fired.
type OnDemoUploadEndedPayload struct {
	Event
	MatchID   string `json:"matchid"`
	MapNumber int    `json:"map_number"`
	Filename  string `json:"filename"`
	Success   bool   `json:"success"`
}

// Map Flow

// OnMatchPausedPayload Fired when the match is paused.
type OnMatchPausedPayload struct {
	Event
	MatchID   string `json:"matchid"`
	MapNumber int    `json:"map_number"`
	Team      string `json:"team"`
	PauseType string `json:"pause_type"`
}

// OnMatchUnpausedPayload Fired when the match is unpaused.
type OnMatchUnpausedPayload struct {
	Event
	MatchID   string `json:"matchid"`
	MapNumber int    `json:"map_number"`
	Team      string `json:"team"`
	PauseType string `json:"pause_type"`
}

type OnPauseBeganPayload struct {
	Event
	MatchID   string `json:"matchid"`
	MapNumber int    `json:"map_number"`
	Team      string `json:"team"`
	PauseType string `json:"pause_type"`
}

// OnKnifeRoundStartedPayload Fired when the knife round starts.
type OnKnifeRoundStartedPayload struct {
	Event
	MatchID   string `json:"matchid"`
	MapNumber int    `json:"map_number"`
}

// OnKnifeRoundWonPayload Fired when the knife round is over and the teams have elected to swap or stay. side represents the chosen side of the winning team, not the side that won the knife round.
type OnKnifeRoundWonPayload struct {
	Event
	MatchID   string `json:"matchid"`
	MapNumber int    `json:"map_number"`
	Team      string `json:"team"`
	Side      string `json:"side"`
	Swapped   bool   `json:"swapped"`
}

// OnTeamReadyStatusChangedPayload Fired when a team's ready status changes.
type OnTeamReadyStatusChangedPayload struct {
	Event
	MatchID   string  `json:"matchid"`
	Team      *string `json:"team"` // nullable
	Ready     bool    `json:"ready"`
	GameState string  `json:"game_state"`
}

// OnGoingLivePayload Fired when a map is going live.
type OnGoingLivePayload struct {
	Event
	MatchID   string `json:"matchid"`
	MapNumber int    `json:"map_number"`
}

// OnRoundStartPayload Fired when a round starts (when freezetime begins).
type OnRoundStartPayload struct {
	Event
	MatchID     string `json:"matchid"`
	MapNumber   int    `json:"map_number"`
	RoundNumber int    `json:"round_number"`
}

// OnRoundEndPayload Fired when a round ends - when the result is in; not when the round stops. Game activity can occur after this.
type OnRoundEndPayload struct {
	Event
	MatchID     string        `json:"matchid"`
	MapNumber   int           `json:"map_number"`
	RoundNumber int           `json:"round_number"`
	RoundTime   int           `json:"round_time"`
	Reason      int           `json:"reason"`
	Winner      Winner        `json:"winner"`
	Team1       Get5StatsTeam `json:"team1"`
	Team2       Get5StatsTeam `json:"team2"`
}

// Winner Winner team
type Winner struct {
	Side string `json:"side"`
	Team string `json:"team"`
}

// OnRoundStatsUpdatedPayload Fired after the stats update on round end.
type OnRoundStatsUpdatedPayload struct {
	Event
	MatchID     string `json:"matchid"`
	MapNumber   int    `json:"map_number"`
	RoundNumber int    `json:"round_number"`
}

// OnPlayerBecameMVPPayload Fired when a player is elected the MVP of the round.
type OnPlayerBecameMVPPayload struct {
	Event
	MatchID     string `json:"matchid"`
	MapNumber   int    `json:"map_number"`
	RoundNumber int    `json:"round_number"`
	Player      Player `json:"player"`
	Reason      int    `json:"reason"`
}

// Player player
type Player struct {
	UserID  int    `json:"user_id"`
	SteamID string `json:"SteamID"`
	Side    string `json:"side"`
	Name    string `json:"name"`
	IsBot   bool   `json:"is_bot"`
}

// OnGrenadeThrownPayload Fired whenever a grenade is thrown by a player. The weapon property reflects the grenade used.
type OnGrenadeThrownPayload struct {
	Event
	MatchID     string `json:"matchid"`
	MapNumber   int    `json:"map_number"`
	RoundNumber int    `json:"round_number"`
	RoundTime   int    `json:"round_time"`
	Player      Player `json:"player"`
	Weapon      Weapon `json:"weapon"`
}

// Weapon Weapon
type Weapon struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// OnPlayerDeathPayload Fired when a player dies.
type OnPlayerDeathPayload struct {
	Event
	MatchID       string    `json:"matchid"`
	MapNumber     int       `json:"map_number"`
	RoundNumber   int       `json:"round_number"`
	RoundTime     int       `json:"round_time"`
	Player        Player    `json:"player"`
	Weapon        Weapon    `json:"weapon"`
	Bomb          bool      `json:"bomb"`
	Headshot      bool      `json:"headshot"`
	ThruSmoke     bool      `json:"thru_smoke"`
	Penetrated    float64   `json:"penetrated"` // Potentially bool
	AttackerBlind bool      `json:"attacker_blind"`
	NoScope       bool      `json:"no_scope"`
	Suicide       bool      `json:"suicide"`
	FriendlyFire  bool      `json:"friendly_fire"`
	Attacker      *Attacker `json:"attacker"` // nullable
	Assist        *Assist   `json:"assist"`   // nullable
}

// Attacker attacker player
type Attacker struct {
	UserID  int    `json:"user_id"`
	SteamID string `json:"SteamID"`
	Side    string `json:"side"`
	Name    string `json:"name"`
	IsBot   bool   `json:"is_bot"`
}

// Assist assister data
type Assist struct {
	Player       Player `json:"player"`
	FriendlyFire bool   `json:"friendly_fire"`
	FlashAssist  bool   `json:"flash_assist"`
}

// OnHEGrenadeDetonatedPayload Fired when an HE grenade detonates. player describes who threw the HE and victims who were affected. weapon is always an HE grenade.
type OnHEGrenadeDetonatedPayload struct {
	Event
	MatchID          string   `json:"matchid"`
	MapNumber        int      `json:"map_number"`
	RoundNumber      int      `json:"round_number"`
	RoundTime        int      `json:"round_time"`
	Player           Player   `json:"player"`
	Weapon           Weapon   `json:"weapon"`
	Victims          []Victim `json:"victims"`
	DamageEnemies    int      `json:"damage_enemies"`
	DamageFriendlies int      `json:"damage_friendlies"`
}

// Victim Victim Player
type Victim struct {
	Player        Player  `json:"player"`
	FriendlyFire  bool    `json:"friendly_fire"`
	BlindDuration float64 `json:"blind_duration,omitempty"`
}

// OnMolotovDetonatedPayload Fired when a molotov grenade expires. player describes who threw the molotov and victims who were affected. weapon is always a molotov grenade. Note that round_time reflects the time at which the grenade detonated (started burning).
type OnMolotovDetonatedPayload struct {
	Event
	MatchID          string   `json:"matchid"`
	MapNumber        int      `json:"map_number"`
	RoundNumber      int      `json:"round_number"`
	RoundTime        int      `json:"round_time"`
	Player           Player   `json:"player"`
	Weapon           Weapon   `json:"weapon"`
	Victims          []Victim `json:"victims"`
	DamageEnemies    int      `json:"damage_enemies"`
	DamageFriendlies int      `json:"damage_friendlies"`
}

// OnFlashbangDetonatedPayload Fired when a flash bang grenade detonates. player describes who threw the flash bang and victims who were affected. weapon is always a flash bang grenade.
type OnFlashbangDetonatedPayload struct {
	Event
	MatchID     string   `json:"matchid"`
	MapNumber   int      `json:"map_number"`
	RoundNumber int      `json:"round_number"`
	RoundTime   int      `json:"round_time"`
	Player      Player   `json:"player"`
	Weapon      Weapon   `json:"weapon"`
	Victims     []Victim `json:"victims"`
}

// OnSmokeGrenadeDetonatedPayload Fired when an smoke grenade expires. player describes who threw the grenade. weapon is always a smoke grenade.
type OnSmokeGrenadeDetonatedPayload struct {
	Event
	MatchID             string `json:"matchid"`
	MapNumber           int    `json:"map_number"`
	RoundNumber         int    `json:"round_number"`
	RoundTime           int    `json:"round_time"`
	Player              Player `json:"player"`
	Weapon              Weapon `json:"weapon"`
	ExtinguishedMolotov bool   `json:"extinguished_molotov"`
}

// OnDecoyStartedPayload Fired when a decoy starts making noise. player describes who threw the grenade. weapon is always a decoy grenade.
type OnDecoyStartedPayload struct {
	Event
	MatchID     string `json:"matchid"`
	MapNumber   int    `json:"map_number"`
	RoundNumber int    `json:"round_number"`
	RoundTime   int    `json:"round_time"`
	Player      Player `json:"player"`
	Weapon      Weapon `json:"weapon"`
}

// OnBombPlantedPayload Fired when the bomb is planted. player describes who planted the bomb.
type OnBombPlantedPayload struct {
	Event
	MatchID     string `json:"matchid"`
	MapNumber   int    `json:"map_number"`
	RoundNumber int    `json:"round_number"`
	RoundTime   int    `json:"round_time"`
	Player      Player `json:"player"`
	Site        string `json:"site"`
}

// OnBombDefusedPayload Fired when the bomb is defused. player describes who defused the bomb.
type OnBombDefusedPayload struct {
	Event
	MatchID           string `json:"matchid"`
	MapNumber         int    `json:"map_number"`
	RoundNumber       int    `json:"round_number"`
	RoundTime         int    `json:"round_time"`
	Player            Player `json:"player"`
	Site              string `json:"site"`
	BombTimeRemaining int    `json:"bomb_time_remaining"`
}

// OnBombExplodedPayload Fired when the bomb explodes.
type OnBombExplodedPayload struct {
	Event
	MatchID     string `json:"matchid"`
	MapNumber   int    `json:"map_number"`
	RoundNumber int    `json:"round_number"`
	RoundTime   int    `json:"round_time"`
	Site        string `json:"site"`
}

// OnPlayerConnectedPayload Fired when a player connects to the server.
type OnPlayerConnectedPayload struct {
	Event
	MatchID   string `json:"matchid"`
	Player    Player `json:"player"`
	IPAddress string `json:"ip_address"`
}

// OnPlayerDisconnectedPayload Fired when a player disconnects from the server.
type OnPlayerDisconnectedPayload struct {
	Event
	MatchID string `json:"matchid"`
	Player  Player `json:"player"`
}

// OnPlayerSayPayload Fired when a player types in chat.
type OnPlayerSayPayload struct {
	Event
	MatchID     string `json:"matchid"`
	MapNumber   int    `json:"map_number"`
	RoundNumber int    `json:"round_number"`
	RoundTime   int    `json:"round_time"`
	Player      Player `json:"player"`
	Command     string `json:"command"`
	Message     string `json:"message"`
}
