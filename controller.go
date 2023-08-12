package got5

import (
	"context"
	"io"
)

// Auth Auth interface handles auth
type Auth interface {
	EventAuth(ctx context.Context, serverID string, auth string) error
	MatchAuth(ctx context.Context, mid string, auth string) error
	CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID string, auth string) error
}

// EventHandler EventHandler interface handles read operation by get5 events
type EventHandler interface {
	Close() error

	// GET5 Events
	HandleOnGameStateChanged(ctx context.Context, p OnGameStateChangedPayload) error
	HandleOnPreLoadMatchConfig(ctx context.Context, p OnPreLoadMatchConfigPayload) error
	HandleOnLoadMatchConfigFailed(ctx context.Context, p OnLoadMatchConfigFailedPayload) error
	HandleOnSeriesInit(ctx context.Context, p OnSeriesInitPayload) error
	HandleOnMapResult(ctx context.Context, p OnMapResultPayload) error
	HandleOnSeriesResult(ctx context.Context, p OnSeriesResultPayload) error
	HandleOnSidePicked(ctx context.Context, p OnSidePickedPayload) error
	HandleOnMapPicked(ctx context.Context, p OnMapPickedPayload) error
	HandleOnMapVetoed(ctx context.Context, p OnMapVetoedPayload) error
	HandleOnBackupRestore(ctx context.Context, p OnBackupRestorePayload) error
	HandleOnDemoFinished(ctx context.Context, p OnDemoFinishedPayload) error
	HandleOnDemoUploadEnded(ctx context.Context, p OnDemoUploadEndedPayload) error
	HandleOnMatchPaused(ctx context.Context, p OnMatchPausedPayload) error
	HandleOnMatchUnpaused(ctx context.Context, p OnMatchUnpausedPayload) error
	HandleOnPauseBegan(ctx context.Context, p OnPauseBeganPayload) error
	HandleOnKnifeRoundStarted(ctx context.Context, p OnKnifeRoundStartedPayload) error
	HandleOnKnifeRoundWon(ctx context.Context, p OnKnifeRoundWonPayload) error
	HandleOnTeamReadyStatusChanged(ctx context.Context, p OnTeamReadyStatusChangedPayload) error
	HandleOnGoingLive(ctx context.Context, p OnGoingLivePayload) error
	HandleOnRoundStart(ctx context.Context, p OnRoundStartPayload) error
	HandleOnRoundEnd(ctx context.Context, p OnRoundEndPayload) error
	HandleOnRoundStatsUpdated(ctx context.Context, p OnRoundStatsUpdatedPayload) error
	HandleOnPlayerBecameMVP(ctx context.Context, p OnPlayerBecameMVPPayload) error
	HandleOnGrenadeThrown(ctx context.Context, p OnGrenadeThrownPayload) error
	HandleOnPlayerDeath(ctx context.Context, p OnPlayerDeathPayload) error
	HandleOnHEGrenadeDetonated(ctx context.Context, p OnHEGrenadeDetonatedPayload) error
	HandleOnMolotovDetonated(ctx context.Context, p OnMolotovDetonatedPayload) error
	HandleOnFlashbangDetonated(ctx context.Context, p OnFlashbangDetonatedPayload) error
	HandleOnSmokeGrenadeDetonated(ctx context.Context, p OnSmokeGrenadeDetonatedPayload) error
	HandleOnDecoyStarted(ctx context.Context, p OnDecoyStartedPayload) error
	HandleOnBombPlanted(ctx context.Context, p OnBombPlantedPayload) error
	HandleOnBombDefused(ctx context.Context, p OnBombDefusedPayload) error
	HandleOnBombExploded(ctx context.Context, p OnBombExplodedPayload) error
	HandleOnPlayerConnected(ctx context.Context, p OnPlayerConnectedPayload) error
	HandleOnPlayerDisconnected(ctx context.Context, p OnPlayerDisconnectedPayload) error
	HandleOnPlayerSay(ctx context.Context, p OnPlayerSayPayload) error
}

// MatchLoader is for Read Operation(get5_loadmatch_url)
type MatchLoader interface {
	// Load respond to get5_loadmatch_url
	Load(ctx context.Context, mid string) (G5Match, error)
}

// DemoUploader is for Demo Upload Operation(get5_dem_upload_url)
type DemoUploader interface {
	Upload(ctx context.Context, mid string, filename string, r io.Reader) error // demoファイルの登録処理
}
