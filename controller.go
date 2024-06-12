package got5

import (
	"context"
	"io"
)

// Auth Auth interface handles auth
type Auth[TMatchID MatchID] interface {
	EventAuth(ctx context.Context, serverID string, auth string) error
	MatchAuth(ctx context.Context, mid TMatchID, auth string) error
	CheckDemoAuth(ctx context.Context, mid TMatchID, filename string, mapNumber int, serverID string, auth string) error
}

// EventHandler EventHandler interface handles read operation by get5 events
type EventHandler[TMatchID MatchID] interface {
	Close() error

	// GET5 Events
	HandleOnGameStateChanged(ctx context.Context, p OnGameStateChangedPayload) error
	HandleOnPreLoadMatchConfig(ctx context.Context, p OnPreLoadMatchConfigPayload) error
	HandleOnLoadMatchConfigFailed(ctx context.Context, p OnLoadMatchConfigFailedPayload) error
	HandleOnSeriesInit(ctx context.Context, p OnSeriesInitPayload[TMatchID]) error
	HandleOnMapResult(ctx context.Context, p OnMapResultPayload[TMatchID]) error
	HandleOnSeriesResult(ctx context.Context, p OnSeriesResultPayload[TMatchID]) error
	HandleOnSidePicked(ctx context.Context, p OnSidePickedPayload[TMatchID]) error
	HandleOnMapPicked(ctx context.Context, p OnMapPickedPayload[TMatchID]) error
	HandleOnMapVetoed(ctx context.Context, p OnMapVetoedPayload[TMatchID]) error
	HandleOnBackupRestore(ctx context.Context, p OnBackupRestorePayload[TMatchID]) error
	HandleOnDemoFinished(ctx context.Context, p OnDemoFinishedPayload[TMatchID]) error
	HandleOnDemoUploadEnded(ctx context.Context, p OnDemoUploadEndedPayload[TMatchID]) error
	HandleOnMatchPaused(ctx context.Context, p OnMatchPausedPayload[TMatchID]) error
	HandleOnMatchUnpaused(ctx context.Context, p OnMatchUnpausedPayload[TMatchID]) error
	HandleOnPauseBegan(ctx context.Context, p OnPauseBeganPayload[TMatchID]) error
	HandleOnKnifeRoundStarted(ctx context.Context, p OnKnifeRoundStartedPayload[TMatchID]) error
	HandleOnKnifeRoundWon(ctx context.Context, p OnKnifeRoundWonPayload[TMatchID]) error
	HandleOnTeamReadyStatusChanged(ctx context.Context, p OnTeamReadyStatusChangedPayload[TMatchID]) error
	HandleOnGoingLive(ctx context.Context, p OnGoingLivePayload[TMatchID]) error
	HandleOnRoundStart(ctx context.Context, p OnRoundStartPayload[TMatchID]) error
	HandleOnRoundEnd(ctx context.Context, p OnRoundEndPayload[TMatchID]) error
	HandleOnRoundStatsUpdated(ctx context.Context, p OnRoundStatsUpdatedPayload[TMatchID]) error
	HandleOnPlayerBecameMVP(ctx context.Context, p OnPlayerBecameMVPPayload[TMatchID]) error
	HandleOnGrenadeThrown(ctx context.Context, p OnGrenadeThrownPayload[TMatchID]) error
	HandleOnPlayerDeath(ctx context.Context, p OnPlayerDeathPayload[TMatchID]) error
	HandleOnHEGrenadeDetonated(ctx context.Context, p OnHEGrenadeDetonatedPayload[TMatchID]) error
	HandleOnMolotovDetonated(ctx context.Context, p OnMolotovDetonatedPayload[TMatchID]) error
	HandleOnFlashbangDetonated(ctx context.Context, p OnFlashbangDetonatedPayload[TMatchID]) error
	HandleOnSmokeGrenadeDetonated(ctx context.Context, p OnSmokeGrenadeDetonatedPayload[TMatchID]) error
	HandleOnDecoyStarted(ctx context.Context, p OnDecoyStartedPayload[TMatchID]) error
	HandleOnBombPlanted(ctx context.Context, p OnBombPlantedPayload[TMatchID]) error
	HandleOnBombDefused(ctx context.Context, p OnBombDefusedPayload[TMatchID]) error
	HandleOnBombExploded(ctx context.Context, p OnBombExplodedPayload[TMatchID]) error
	HandleOnPlayerConnected(ctx context.Context, p OnPlayerConnectedPayload[TMatchID]) error
	HandleOnPlayerDisconnected(ctx context.Context, p OnPlayerDisconnectedPayload[TMatchID]) error
	HandleOnPlayerSay(ctx context.Context, p OnPlayerSayPayload[TMatchID]) error
}

// MatchLoader is for Read Operation(get5_loadmatch_url)
type MatchLoader[TMatchID MatchID] interface {
	// Load respond to get5_loadmatch_url
	Load(ctx context.Context, mid TMatchID) (G5Match[TMatchID], error)
}

// DemoUploader is for Demo Upload Operation(get5_dem_upload_url)
type DemoUploader[TMatchID MatchID] interface {
	Upload(ctx context.Context, mid TMatchID, filename string, r io.Reader) error // demoファイルの登録処理
}
