package logger

import (
	"context"
	"log"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

var _ controller.EventHandler = (*loggerEventHandler)(nil)

// logger is simply prints what happend.
type loggerEventHandler struct{}

// HandleOnPauseBegan implements controller.EventHandler.
func (*loggerEventHandler) HandleOnPauseBegan(ctx context.Context, p models.OnPauseBeganPayload) error {
	log.Println(p)
	return nil
}

// CheckAuth implements controller.EventHandler
func (*loggerEventHandler) CheckEventAuth(ctx context.Context, mid string, reqAuth string) error {
	return nil
}

// Hostname implements controller.EventHandler
func (*loggerEventHandler) Hostname() string {
	return ""
}

func (*loggerEventHandler) Close() error {
	return nil
}

// HandleOnBackupRestore implements controller.EventHandler
func (*loggerEventHandler) HandleOnBackupRestore(ctx context.Context, p models.OnBackupRestorePayload) error {
	log.Println(p)
	return nil
}

// HandleOnBombDefused implements controller.EventHandler
func (*loggerEventHandler) HandleOnBombDefused(ctx context.Context, p models.OnBombDefusedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnBombExploded implements controller.EventHandler
func (*loggerEventHandler) HandleOnBombExploded(ctx context.Context, p models.OnBombExplodedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnBombPlanted implements controller.EventHandler
func (*loggerEventHandler) HandleOnBombPlanted(ctx context.Context, p models.OnBombPlantedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnDecoyStarted implements controller.EventHandler
func (*loggerEventHandler) HandleOnDecoyStarted(ctx context.Context, p models.OnDecoyStartedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnDemoFinished implements controller.EventHandler
func (*loggerEventHandler) HandleOnDemoFinished(ctx context.Context, p models.OnDemoFinishedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnDemoUploadEnded implements controller.EventHandler
func (*loggerEventHandler) HandleOnDemoUploadEnded(ctx context.Context, p models.OnDemoUploadEndedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnEvent implements controller.EventHandler
func (*loggerEventHandler) HandleOnEvent(ctx context.Context, p models.OnEventPayload) error {
	log.Println(p)
	return nil
}

// HandleOnFlashbangDetonated implements controller.EventHandler
func (*loggerEventHandler) HandleOnFlashbangDetonated(ctx context.Context, p models.OnFlashbangDetonatedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnGameStateChanged implements controller.EventHandler
func (*loggerEventHandler) HandleOnGameStateChanged(ctx context.Context, p models.OnGameStateChangedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnGoingLive implements controller.EventHandler
func (*loggerEventHandler) HandleOnGoingLive(ctx context.Context, p models.OnGoingLivePayload) error {
	log.Println(p)
	return nil
}

// HandleOnGrenadeThrown implements controller.EventHandler
func (*loggerEventHandler) HandleOnGrenadeThrown(ctx context.Context, p models.OnGrenadeThrownPayload) error {
	log.Println(p)
	return nil
}

// HandleOnHEGrenadeDetonated implements controller.EventHandler
func (*loggerEventHandler) HandleOnHEGrenadeDetonated(ctx context.Context, p models.OnHEGrenadeDetonatedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnKnifeRoundStarted implements controller.EventHandler
func (*loggerEventHandler) HandleOnKnifeRoundStarted(ctx context.Context, p models.OnKnifeRoundStartedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnKnifeRoundWon implements controller.EventHandler
func (*loggerEventHandler) HandleOnKnifeRoundWon(ctx context.Context, p models.OnKnifeRoundWonPayload) error {
	log.Println(p)
	return nil
}

// HandleOnLoadMatchConfigFailed implements controller.EventHandler
func (*loggerEventHandler) HandleOnLoadMatchConfigFailed(ctx context.Context, p models.OnLoadMatchConfigFailedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnMapPicked implements controller.EventHandler
func (*loggerEventHandler) HandleOnMapPicked(ctx context.Context, p models.OnMapPickedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnMapResult implements controller.EventHandler
func (*loggerEventHandler) HandleOnMapResult(ctx context.Context, p models.OnMapResultPayload) error {
	log.Println(p)
	return nil
}

// HandleOnMapVetoed implements controller.EventHandler
func (*loggerEventHandler) HandleOnMapVetoed(ctx context.Context, p models.OnMapVetoedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnMatchPaused implements controller.EventHandler
func (*loggerEventHandler) HandleOnMatchPaused(ctx context.Context, p models.OnMatchPausedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnMatchUnpaused implements controller.EventHandler
func (*loggerEventHandler) HandleOnMatchUnpaused(ctx context.Context, p models.OnMatchUnpausedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnMolotovDetonated implements controller.EventHandler
func (*loggerEventHandler) HandleOnMolotovDetonated(ctx context.Context, p models.OnMolotovDetonatedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnPlayerBecameMVP implements controller.EventHandler
func (*loggerEventHandler) HandleOnPlayerBecameMVP(ctx context.Context, p models.OnPlayerBecameMVPPayload) error {
	log.Println(p)
	return nil
}

// HandleOnPlayerConnected implements controller.EventHandler
func (*loggerEventHandler) HandleOnPlayerConnected(ctx context.Context, p models.OnPlayerConnectedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnPlayerDeath implements controller.EventHandler
func (*loggerEventHandler) HandleOnPlayerDeath(ctx context.Context, p models.OnPlayerDeathPayload) error {
	log.Println(p)
	return nil
}

// HandleOnPlayerDisconnected implements controller.EventHandler
func (*loggerEventHandler) HandleOnPlayerDisconnected(ctx context.Context, p models.OnPlayerDisconnectedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnPlayerSay implements controller.EventHandler
func (*loggerEventHandler) HandleOnPlayerSay(ctx context.Context, p models.OnPlayerSayPayload) error {
	log.Println(p)
	return nil
}

// HandleOnPreLoadMatchConfig implements controller.EventHandler
func (*loggerEventHandler) HandleOnPreLoadMatchConfig(ctx context.Context, p models.OnPreLoadMatchConfigPayload) error {
	log.Println(p)
	return nil
}

// HandleOnRoundEnd implements controller.EventHandler
func (*loggerEventHandler) HandleOnRoundEnd(ctx context.Context, p models.OnRoundEndPayload) error {
	log.Println(p)
	return nil
}

// HandleOnRoundStart implements controller.EventHandler
func (*loggerEventHandler) HandleOnRoundStart(ctx context.Context, p models.OnRoundStartPayload) error {
	log.Println(p)
	return nil
}

// HandleOnRoundStatsUpdated implements controller.EventHandler
func (*loggerEventHandler) HandleOnRoundStatsUpdated(ctx context.Context, p models.OnRoundStatsUpdatedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnSeriesInit implements controller.EventHandler
func (*loggerEventHandler) HandleOnSeriesInit(ctx context.Context, p models.OnSeriesInitPayload) error {
	log.Println(p)
	return nil
}

// HandleOnSeriesResult implements controller.EventHandler
func (*loggerEventHandler) HandleOnSeriesResult(ctx context.Context, p models.OnSeriesResultPayload) error {
	log.Println(p)
	return nil
}

// HandleOnSidePicked implements controller.EventHandler
func (*loggerEventHandler) HandleOnSidePicked(ctx context.Context, p models.OnSidePickedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnSmokeGrenadeDetonated implements controller.EventHandler
func (*loggerEventHandler) HandleOnSmokeGrenadeDetonated(ctx context.Context, p models.OnSmokeGrenadeDetonatedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnTeamReadyStatusChanged implements controller.EventHandler
func (*loggerEventHandler) HandleOnTeamReadyStatusChanged(ctx context.Context, p models.OnTeamReadyStatusChangedPayload) error {
	log.Println(p)
	return nil
}

// NewEventHandler Get Logger pointer
func NewEventHandler() controller.EventHandler {
	return &loggerEventHandler{}
}
