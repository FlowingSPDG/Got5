# Got5
[![test](https://github.com/FlowingSPDG/Got5/actions/workflows/dagger.yml/badge.svg)](https://github.com/FlowingSPDG/Got5/actions/workflows/dagger.yml)  
STATUS: UNDER DEVELOPMENT  
Go + get5 = Got5!  
Got5 is simple and fast, Build-Your-Own-Get5Web framework.  
Built with Go(1.20), Fiber.  

Got5 interfaces supports general get5 event handling such as Kill, Match Load or Demo upload.  
This makes you easier to build get5-based system like get5-web.  

## Developer
Shugo Kawamura - 河村 柊吾

### Got5 interfaces
[controller](https://github.com/FlowingSPDG/Got5/tree/main/controller) package has 3 interfaces that communicates with get5.  
controller does not have database interface, so you may implement database system by yourself.  

#### EventHandler
[EventHandler](https://github.com/FlowingSPDG/Got5/blob/75996d44058558ca7453af1c4b4f9e73115924d4/controller/controller.go#L10-L52) interface should handle event coming from get5 [Events & Forwards](https://splewis.github.io/get5/latest/events_and_forwards/) (get5_remote_log_url).  
e.g. You can post Discord message, or save stats to your database.  
```go
// EventHandler EventHandler interface handles read operation by get5 events
type EventHandler interface {
	Close() error

	// Auth Checker
	CheckEventAuth(ctx context.Context, mid string, reqAuth string) error

	// GET5 Events
	HandleOnGameStateChanged(ctx context.Context, p models.OnGameStateChangedPayload) error
	HandleOnPreLoadMatchConfig(ctx context.Context, p models.OnPreLoadMatchConfigPayload) error
	HandleOnLoadMatchConfigFailed(ctx context.Context, p models.OnLoadMatchConfigFailedPayload) error
	HandleOnSeriesInit(ctx context.Context, p models.OnSeriesInitPayload) error
	HandleOnMapResult(ctx context.Context, p models.OnMapResultPayload) error
	HandleOnSeriesResult(ctx context.Context, p models.OnSeriesResultPayload) error
	HandleOnSidePicked(ctx context.Context, p models.OnSidePickedPayload) error
	HandleOnMapPicked(ctx context.Context, p models.OnMapPickedPayload) error
	HandleOnMapVetoed(ctx context.Context, p models.OnMapVetoedPayload) error
	HandleOnBackupRestore(ctx context.Context, p models.OnBackupRestorePayload) error
	HandleOnDemoFinished(ctx context.Context, p models.OnDemoFinishedPayload) error
	HandleOnDemoUploadEnded(ctx context.Context, p models.OnDemoUploadEndedPayload) error
	HandleOnMatchPaused(ctx context.Context, p models.OnMatchPausedPayload) error
	HandleOnMatchUnpaused(ctx context.Context, p models.OnMatchUnpausedPayload) error
	HandleOnKnifeRoundStarted(ctx context.Context, p models.OnKnifeRoundStartedPayload) error
	HandleOnKnifeRoundWon(ctx context.Context, p models.OnKnifeRoundWonPayload) error
	HandleOnTeamReadyStatusChanged(ctx context.Context, p models.OnTeamReadyStatusChangedPayload) error
	HandleOnGoingLive(ctx context.Context, p models.OnGoingLivePayload) error
	HandleOnRoundStart(ctx context.Context, p models.OnRoundStartPayload) error
	HandleOnRoundEnd(ctx context.Context, p models.OnRoundEndPayload) error
	HandleOnRoundStatsUpdated(ctx context.Context, p models.OnRoundStatsUpdatedPayload) error
	HandleOnPlayerBecameMVP(ctx context.Context, p models.OnPlayerBecameMVPPayload) error
	HandleOnGrenadeThrown(ctx context.Context, p models.OnGrenadeThrownPayload) error
	HandleOnPlayerDeath(ctx context.Context, p models.OnPlayerDeathPayload) error
	HandleOnHEGrenadeDetonated(ctx context.Context, p models.OnHEGrenadeDetonatedPayload) error
	HandleOnMolotovDetonated(ctx context.Context, p models.OnMolotovDetonatedPayload) error
	HandleOnFlashbangDetonated(ctx context.Context, p models.OnFlashbangDetonatedPayload) error
	HandleOnSmokeGrenadeDetonated(ctx context.Context, p models.OnSmokeGrenadeDetonatedPayload) error
	HandleOnDecoyStarted(ctx context.Context, p models.OnDecoyStartedPayload) error
	HandleOnBombPlanted(ctx context.Context, p models.OnBombPlantedPayload) error
	HandleOnBombDefused(ctx context.Context, p models.OnBombDefusedPayload) error
	HandleOnBombExploded(ctx context.Context, p models.OnBombExplodedPayload) error
	HandleOnPlayerConnected(ctx context.Context, p models.OnPlayerConnectedPayload) error
	HandleOnPlayerDisconnected(ctx context.Context, p models.OnPlayerDisconnectedPayload) error
	HandleOnPlayerSay(ctx context.Context, p models.OnPlayerSayPayload) error
}
```

#### MatchLoader
[MatchLoader](https://github.com/FlowingSPDG/Got5/blob/75996d44058558ca7453af1c4b4f9e73115924d4/controller/controller.go#L54-L57) interface should handle ``get5_loadmatch_url`` request from game server.  
You need to respond game server with JSON.  
https://splewis.github.io/get5/latest/match_schema/  
```go
// MatchLoader is for Read Operation(get5_loadmatch_url)
type MatchLoader interface {
	// Auth Checker
	CheckMatchAuth(ctx context.Context, mid string, auth string) error

	// Load respond to get5_loadmatch_url
	Load(ctx context.Context, mid string) (models.G5Match, error)
}
```

models.G5Match is interface for generating get5 supported format JSON.  
```go
// G5Match 別の構造体にG5Matchインターフェースを実装すれば型が違っても変換してGet5に渡してくれる
type G5Match interface {
	ToG5Format() Match
}
```
```go
type Match struct {
	MatchTitle           string            `json:"match_title"`
	MatchID              string            `json:"matchid"`
	ClinchSeries         bool              `json:"clinch_series"`
	NumMaps              int               `json:"num_maps"`
	PlayersPerTeam       int               `json:"players_per_team"`
	CoachesPerTeam       int               `json:"coaches_per_team"`
	CoachesMustReady     bool              `json:"coaches_must_ready"`
	MinPlayersToReady    int               `json:"min_players_to_ready"`
	MinSpectatorsToReady int               `json:"min_spectators_to_ready"`
	SkipVeto             bool              `json:"skip_veto"`
	VetoFirst            string            `json:"veto_first"`
	SideType             string            `json:"side_type"`
	Spectators           Spectators        `json:"spectators"`
	Maplist              []string          `json:"maplist"`
	MapSides             []string          `json:"map_sides"`
	Team1                Team              `json:"team1"`
	Team2                Team              `json:"team2"`
	Cvars                map[string]string `json:"cvars"`
}
```

#### DemoUploader
[DemoUploader](https://github.com/FlowingSPDG/Got5/blob/75996d44058558ca7453af1c4b4f9e73115924d4/controller/controller.go#L59-L62) interface should handle demo upload from game server.  
You may want to add auth middleware to prevend unauthorized demo uploads.  
```go
// DemoUploader is for Demo Upload Operation(get5_dem_upload_url)
type DemoUploader interface {
	CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID int, auth string) error
	Upload(ctx context.Context, mid string, filename string, r io.Reader) error // demoファイルの登録処理
}
```

## Authentication  
There are some CVARs for authenticating Got5 and get5 gameserver.  
- `get5_remote_log_url` - For specifying where to send Event Data.  
  - `get5_remote_log_header_key` - HTTP Header key for event request.  
  - `get5_remote_log_header_value` - HTTP header value for event request.  
- `get5_demo_upload_url` - For specifying where to upload demo data.  
  - `get5_demo_upload_header_key` - HTTP Header key for demo upload.  
  - `get5_demo_upload_header_value` - HTTP Header value for demo upload.  
- `get5_loadmatch_url "https://example.com/match_config.json" "Authorization" "Bearer <token>"` - For loading match info.  

## Examples
### logger
[logger](https://github.com/FlowingSPDG/Got5/tree/main/examples/discord) is most simple implemention for Got5 interfaces.  
logger simply prints what happens.  
No write operation, No Data store system.  

### firebase
[firebase](https://github.com/FlowingSPDG/Got5/tree/main/examples/firebase) interfaces will read/write match/event informations.  
You may need to enable firestore on your Google Cloud Platform project.  

### Discord
[discord](https://github.com/FlowingSPDG/Got5/tree/main/examples/discord) is most complicated implemention of Got5 interfaces.  
Discord struct implements all Got5 interfaces.  
You can create match, store event datas and post what happens.  

### Discord webhook
[discord_webhook](https://github.com/FlowingSPDG/Got5/tree/main/examples/discord_webhook) posts received get5 events.  
This is fairly simple and nice example for learning your first Got5 based system.

