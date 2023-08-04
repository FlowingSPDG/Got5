package controller

import (
	"context"
	"io"

	"github.com/FlowingSPDG/Got5/models"
)

// Auth Auth interface handles auth
type Auth interface {
	EventAuth(ctx context.Context, serverID string, auth string) error
	MatchAuth(ctx context.Context, mid string, auth string) error
	CheckDemoAuth(ctx context.Context, mid string, filename string, mapNumber int, serverID int, auth string) error
}

// EventHandler EventHandler interface handles read operation by get5 events
type EventHandler interface {
	Close() error

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
	HandleOnPauseBegan(ctx context.Context, p models.OnPauseBeganPayload) error
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

// MatchLoader is for Read Operation(get5_loadmatch_url)
type MatchLoader interface {
	// Load respond to get5_loadmatch_url
	Load(ctx context.Context, mid string) (models.G5Match, error)
}

// DemoUploader is for Demo Upload Operation(get5_dem_upload_url)
type DemoUploader interface {
	Upload(ctx context.Context, mid string, filename string, r io.Reader) error // demoファイルの登録処理
}
