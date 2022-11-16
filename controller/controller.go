package controller

import (
	"context"

	"github.com/FlowingSPDG/Got5/models"
)

// Controller Controller interface operates CRUD operation for e.g. Databases.
type Controller interface {
	// TODO: イベントハンドラーの部分のinterfaceとし、CRUDオペレーションに関連したオペレーション部分を別のinterfaceとして再定義する
	Close() error

	// イベントハンドラへの絶対パスURL
	Hostname() string

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

	// Read Operation
	GetMatch(ctx context.Context, mid string) (models.G5Match, error)
	// GetMap etc

	// Create Operation
	RegisterMatch(ctx context.Context, m models.Match) (models.Match, error)                          // 外部からマッチ作成リクエストが発生した際に実行されるハンドラ
	RegisterDemoFile(ctx context.Context, bucket string, mid string, filename string, b []byte) error // demoファイルの登録処理
}
