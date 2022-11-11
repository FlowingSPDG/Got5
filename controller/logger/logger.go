package logger

import (
	"context"
	"log"

	"github.com/FlowingSPDG/Got5/controller"
	"github.com/FlowingSPDG/Got5/models"
)

// logger is simply prints what happend.
type logger struct{}

// HandleOnBackupRestore implements controller.Controller
func (*logger) HandleOnBackupRestore(ctx context.Context, p models.OnBackupRestorePayload) error {
	log.Println(p)
	return nil
}

// HandleOnBombDefused implements controller.Controller
func (*logger) HandleOnBombDefused(ctx context.Context, p models.OnBombDefusedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnBombExploded implements controller.Controller
func (*logger) HandleOnBombExploded(ctx context.Context, p models.OnBombExplodedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnBombPlanted implements controller.Controller
func (*logger) HandleOnBombPlanted(ctx context.Context, p models.OnBombPlantedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnDecoyStarted implements controller.Controller
func (*logger) HandleOnDecoyStarted(ctx context.Context, p models.OnDecoyStartedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnDemoFinished implements controller.Controller
func (*logger) HandleOnDemoFinished(ctx context.Context, p models.OnDemoFinishedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnDemoUploadEnded implements controller.Controller
func (*logger) HandleOnDemoUploadEnded(ctx context.Context, p models.OnDemoUploadEndedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnEvent implements controller.Controller
func (*logger) HandleOnEvent(ctx context.Context, p models.OnEventPayload) error {
	log.Println(p)
	return nil
}

// HandleOnFlashbangDetonated implements controller.Controller
func (*logger) HandleOnFlashbangDetonated(ctx context.Context, p models.OnFlashbangDetonatedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnGameStateChanged implements controller.Controller
func (*logger) HandleOnGameStateChanged(ctx context.Context, p models.OnGameStateChangedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnGoingLive implements controller.Controller
func (*logger) HandleOnGoingLive(ctx context.Context, p models.OnGoingLivePayload) error {
	log.Println(p)
	return nil
}

// HandleOnGrenadeThrown implements controller.Controller
func (*logger) HandleOnGrenadeThrown(ctx context.Context, p models.OnGrenadeThrownPayload) error {
	log.Println(p)
	return nil
}

// HandleOnHEGrenadeDetonated implements controller.Controller
func (*logger) HandleOnHEGrenadeDetonated(ctx context.Context, p models.OnHEGrenadeDetonatedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnKnifeRoundStarted implements controller.Controller
func (*logger) HandleOnKnifeRoundStarted(ctx context.Context, p models.OnKnifeRoundStartedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnKnifeRoundWon implements controller.Controller
func (*logger) HandleOnKnifeRoundWon(ctx context.Context, p models.OnKnifeRoundWonPayload) error {
	log.Println(p)
	return nil
}

// HandleOnLoadMatchConfigFailed implements controller.Controller
func (*logger) HandleOnLoadMatchConfigFailed(ctx context.Context, p models.OnLoadMatchConfigFailedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnMapPicked implements controller.Controller
func (*logger) HandleOnMapPicked(ctx context.Context, p models.OnMapPickedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnMapResult implements controller.Controller
func (*logger) HandleOnMapResult(ctx context.Context, p models.OnMapResultPayload) error {
	log.Println(p)
	return nil
}

// HandleOnMapVetoed implements controller.Controller
func (*logger) HandleOnMapVetoed(ctx context.Context, p models.OnMapVetoedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnMatchPaused implements controller.Controller
func (*logger) HandleOnMatchPaused(ctx context.Context, p models.OnMatchPausedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnMatchUnpaused implements controller.Controller
func (*logger) HandleOnMatchUnpaused(ctx context.Context, p models.OnMatchUnpausedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnMolotovDetonated implements controller.Controller
func (*logger) HandleOnMolotovDetonated(ctx context.Context, p models.OnMolotovDetonatedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnPlayerBecameMVP implements controller.Controller
func (*logger) HandleOnPlayerBecameMVP(ctx context.Context, p models.OnPlayerBecameMVPPayload) error {
	log.Println(p)
	return nil
}

// HandleOnPlayerConnected implements controller.Controller
func (*logger) HandleOnPlayerConnected(ctx context.Context, p models.OnPlayerConnectedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnPlayerDeath implements controller.Controller
func (*logger) HandleOnPlayerDeath(ctx context.Context, p models.OnPlayerDeathPayload) error {
	log.Println(p)
	return nil
}

// HandleOnPlayerDisconnected implements controller.Controller
func (*logger) HandleOnPlayerDisconnected(ctx context.Context, p models.OnPlayerDisconnectedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnPlayerSay implements controller.Controller
func (*logger) HandleOnPlayerSay(ctx context.Context, p models.OnPlayerSayPayload) error {
	log.Println(p)
	return nil
}

// HandleOnPreLoadMatchConfig implements controller.Controller
func (*logger) HandleOnPreLoadMatchConfig(ctx context.Context, p models.OnPreLoadMatchConfigPayload) error {
	log.Println(p)
	return nil
}

// HandleOnRoundEnd implements controller.Controller
func (*logger) HandleOnRoundEnd(ctx context.Context, p models.OnRoundEndPayload) error {
	log.Println(p)
	return nil
}

// HandleOnRoundStart implements controller.Controller
func (*logger) HandleOnRoundStart(ctx context.Context, p models.OnRoundStartPayload) error {
	log.Println(p)
	return nil
}

// HandleOnRoundStatsUpdated implements controller.Controller
func (*logger) HandleOnRoundStatsUpdated(ctx context.Context, p models.OnRoundStatsUpdatedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnSeriesInit implements controller.Controller
func (*logger) HandleOnSeriesInit(ctx context.Context, p models.OnSeriesInitPayload) error {
	log.Println(p)
	return nil
}

// HandleOnSeriesResult implements controller.Controller
func (*logger) HandleOnSeriesResult(ctx context.Context, p models.OnSeriesResultPayload) error {
	log.Println(p)
	return nil
}

// HandleOnSidePicked implements controller.Controller
func (*logger) HandleOnSidePicked(ctx context.Context, p models.OnSidePickedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnSmokeGrenadeDetonated implements controller.Controller
func (*logger) HandleOnSmokeGrenadeDetonated(ctx context.Context, p models.OnSmokeGrenadeDetonatedPayload) error {
	log.Println(p)
	return nil
}

// HandleOnTeamReadyStatusChanged implements controller.Controller
func (*logger) HandleOnTeamReadyStatusChanged(ctx context.Context, p models.OnTeamReadyStatusChangedPayload) error {
	log.Println(p)
	return nil
}

// NewLoggerController Get Logger pointer
func NewLoggerController() controller.Controller {
	return &logger{}
}
