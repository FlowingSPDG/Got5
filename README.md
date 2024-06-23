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
[EventHandler](https://github.com/FlowingSPDG/Got5/blob/75996d44058558ca7453af1c4b4f9e73115924d4/controller/got5.go#L10-L52) interface should handle event coming from get5 [Events & Forwards](https://splewis.github.io/get5/latest/events_and_forwards/) (get5_remote_log_url).  
e.g. You can post Discord message, or save stats to your database.  
```go
// EventHandler EventHandler interface handles read operation by get5 events
type EventHandler interface {
	Close() error

	// Auth Checker
	CheckEventAuth(ctx context.Context, mid string, reqAuth string) error

	// GET5 Events
	HandleOnGameStateChanged(ctx context.Context, p got5.OnGameStateChangedPayload) error
	HandleOnPreLoadMatchConfig(ctx context.Context, p got5.OnPreLoadMatchConfigPayload) error
	HandleOnLoadMatchConfigFailed(ctx context.Context, p got5.OnLoadMatchConfigFailedPayload) error
	HandleOnSeriesInit(ctx context.Context, p got5.OnSeriesInitPayload) error
	HandleOnMapResult(ctx context.Context, p got5.OnMapResultPayload) error
	HandleOnSeriesResult(ctx context.Context, p got5.OnSeriesResultPayload) error
	HandleOnSidePicked(ctx context.Context, p got5.OnSidePickedPayload) error
	HandleOnMapPicked(ctx context.Context, p got5.OnMapPickedPayload) error
	HandleOnMapVetoed(ctx context.Context, p got5.OnMapVetoedPayload) error
	HandleOnBackupRestore(ctx context.Context, p got5.OnBackupRestorePayload) error
	HandleOnDemoFinished(ctx context.Context, p got5.OnDemoFinishedPayload) error
	HandleOnDemoUploadEnded(ctx context.Context, p got5.OnDemoUploadEndedPayload) error
	HandleOnMatchPaused(ctx context.Context, p got5.OnMatchPausedPayload) error
	HandleOnMatchUnpaused(ctx context.Context, p got5.OnMatchUnpausedPayload) error
	HandleOnKnifeRoundStarted(ctx context.Context, p got5.OnKnifeRoundStartedPayload) error
	HandleOnKnifeRoundWon(ctx context.Context, p got5.OnKnifeRoundWonPayload) error
	HandleOnTeamReadyStatusChanged(ctx context.Context, p got5.OnTeamReadyStatusChangedPayload) error
	HandleOnGoingLive(ctx context.Context, p got5.OnGoingLivePayload) error
	HandleOnRoundStart(ctx context.Context, p got5.OnRoundStartPayload) error
	HandleOnRoundEnd(ctx context.Context, p got5.OnRoundEndPayload) error
	HandleOnRoundStatsUpdated(ctx context.Context, p got5.OnRoundStatsUpdatedPayload) error
	HandleOnPlayerBecameMVP(ctx context.Context, p got5.OnPlayerBecameMVPPayload) error
	HandleOnGrenadeThrown(ctx context.Context, p got5.OnGrenadeThrownPayload) error
	HandleOnPlayerDeath(ctx context.Context, p got5.OnPlayerDeathPayload) error
	HandleOnHEGrenadeDetonated(ctx context.Context, p got5.OnHEGrenadeDetonatedPayload) error
	HandleOnMolotovDetonated(ctx context.Context, p got5.OnMolotovDetonatedPayload) error
	HandleOnFlashbangDetonated(ctx context.Context, p got5.OnFlashbangDetonatedPayload) error
	HandleOnSmokeGrenadeDetonated(ctx context.Context, p got5.OnSmokeGrenadeDetonatedPayload) error
	HandleOnDecoyStarted(ctx context.Context, p got5.OnDecoyStartedPayload) error
	HandleOnBombPlanted(ctx context.Context, p got5.OnBombPlantedPayload) error
	HandleOnBombDefused(ctx context.Context, p got5.OnBombDefusedPayload) error
	HandleOnBombExploded(ctx context.Context, p got5.OnBombExplodedPayload) error
	HandleOnPlayerConnected(ctx context.Context, p got5.OnPlayerConnectedPayload) error
	HandleOnPlayerDisconnected(ctx context.Context, p got5.OnPlayerDisconnectedPayload) error
	HandleOnPlayerSay(ctx context.Context, p got5.OnPlayerSayPayload) error
}
```

#### MatchLoader
[MatchLoader](https://github.com/FlowingSPDG/Got5/blob/75996d44058558ca7453af1c4b4f9e73115924d4/controller/got5.go#L54-L57) interface should handle ``get5_loadmatch_url`` request from game server.  
You need to respond JSON.  
https://splewis.github.io/get5/latest/match_schema/  
```go
// MatchLoader is for Read Operation(get5_loadmatch_url)
type MatchLoader interface {
	// Auth Checker
	CheckMatchAuth(ctx context.Context, mid int, auth string) error

	// Load respond to get5_loadmatch_url
	Load(ctx context.Context, mid string) (got5.Match, error)
}
```


#### DemoUploader
[DemoUploader](https://github.com/FlowingSPDG/Got5/blob/75996d44058558ca7453af1c4b4f9e73115924d4/controller/got5.go#L59-L62) interface should handle demo upload from game server.  
You may want to add auth middleware to prevend unauthorized demo uploads.  
```go
// DemoUploader is for Demo Upload Operation(get5_dem_upload_url)
type DemoUploader interface {
	CheckDemoAuth(ctx context.Context, mid int, filename string, mapNumber int, serverID string, auth string) error
	Upload(ctx context.Context, mid int, filename string, r io.Reader) error // demoファイルの登録処理
}
```

## Authentication  
TODO...

## Examples
TODO...